package cluster

import (
	"context"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"reflect"
	"regexp"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	kubeadmv1alpha1 "github.com/iljaweis/kubeadm-operator/pkg/apis/kubeadm/v1alpha1"
	resourcesv1alpha1 "github.com/iljaweis/resource-ctlr/pkg/apis/resources/v1alpha1"
	pkgerrors "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_cluster")

// Add creates a new Cluster Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileCluster{
		client:   mgr.GetClient(),
		scheme:   mgr.GetScheme(),
		recorder: mgr.GetRecorder("kubeadm-operator"),
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("cluster-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Cluster
	err = c.Watch(&source.Kind{Type: &kubeadmv1alpha1.Cluster{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	if err := c.Watch(&source.Kind{Type: &resourcesv1alpha1.Command{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &kubeadmv1alpha1.Cluster{},
	}); err != nil {
		return err
	}

	if err := c.Watch(&source.Kind{Type: &resourcesv1alpha1.File{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &kubeadmv1alpha1.Cluster{},
	}); err != nil {
		return err
	}

	if err := c.Watch(&source.Kind{Type: &resourcesv1alpha1.FileContent{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &kubeadmv1alpha1.Cluster{},
	}); err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileCluster{}

// ReconcileCluster reconciles a Cluster object
type ReconcileCluster struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client   client.Client
	scheme   *runtime.Scheme
	recorder record.EventRecorder
}

// Reconcile reads that state of the cluster for a Cluster object and makes changes based on the state read
// and what is in the Cluster.Spec
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileCluster) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling Cluster")

	cluster := &kubeadmv1alpha1.Cluster{}
	err := r.client.Get(context.TODO(), request.NamespacedName, cluster)
	if err != nil {
		if errors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	update := false

	if err := r.ensureHosts(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "error ensuring target host resources")
	}

	if err := r.ensureKubeadmConfig(logger, cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not create kubeadm config")
	}

	if err := r.runKubeadmOnFirstController(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not run initial kubeadm")
	}

	if err := r.ensureAdminConf(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not fetch and store admin.conf")
	}

	if err := r.ensurePKIStatus(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not set PKI status")
	}

	if err := r.ensureTokenStatus(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not set token status")
	}

	if err := r.ensureNetworking(cluster, &update); err != nil {
		return reconcile.Result{}, pkgerrors.Wrap(err, "could not set up networking")
	}

	if update {
		fmt.Println("update is true")
		logger.Info("updating Cluster status")
		if err := r.client.Status().Update(context.TODO(), cluster); err != nil {
			return reconcile.Result{}, pkgerrors.Wrapf(err, "could not update cluster status for %s", request)
		}

		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}

func (r *ReconcileCluster) ensureHosts(cluster *kubeadmv1alpha1.Cluster, update *bool) error {

	for _, co := range cluster.Spec.Controllers {
		if err := r.ensureHost(cluster, co, update); err != nil {
			return err
		}
	}

	for _, wo := range cluster.Spec.Workers {
		if err := r.ensureHost(cluster, wo, update); err != nil {
			return err
		}
	}

	return nil
}

func (r *ReconcileCluster) ensureHost(cluster *kubeadmv1alpha1.Cluster, chost kubeadmv1alpha1.ClusterMember, update *bool) error {

	host := &resourcesv1alpha1.Host{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: cluster.Namespace, Name: chost.Name}, host)
	if err != nil && !errors.IsNotFound(err) {
		return pkgerrors.Wrapf(err, "could not get host object for %s", chost.Name)
	}

	if err != nil { // not found
		host := &resourcesv1alpha1.Host{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: cluster.Namespace,
				Name:      chost.Name,
			},
			Spec: resourcesv1alpha1.HostSpec{
				IPAddress:    chost.IP,
				SshKeySecret: cluster.Spec.DefaultSSHKeySecret,
				Port:         22,
			},
		}

		_ = controllerutil.SetControllerReference(cluster, host, r.scheme)

		if err := r.client.Create(context.TODO(), host); err != nil {
			return pkgerrors.Wrapf(err, "could not create host object for %s", chost.Name)
		}

		return nil
	}

	return nil
}

func nameForKubeadmConfigFile(cluster *kubeadmv1alpha1.Cluster) string {
	return cluster.Name + "-kubeadm-config"
}

func (r *ReconcileCluster) ensureKubeadmConfig(logger logr.Logger, cluster *kubeadmv1alpha1.Cluster, update *bool) error {

	name := nameForKubeadmConfigFile(cluster)

	kubeadmConfig := &kubeadmv1alpha1.ClusterConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterConfiguration",
			APIVersion: "kubeadm.k8s.io/v1beta1",
		},
		ControlPlaneEndpoint: cluster.Spec.Controllers[0].IP,
	}

	if !reflect.DeepEqual(kubeadmConfig, cluster.Status.KubeadmConfig) {
		cluster.Status.KubeadmConfig = kubeadmConfig
		*update = true
	}

	var newConfig string

	if b, err := yaml.Marshal(cluster.Status.KubeadmConfig); err != nil {
		return pkgerrors.Wrapf(err, "could not encode kubeadm-config.yaml")
	} else {
		newConfig = string(b)
	}

	file := &resourcesv1alpha1.File{ObjectMeta: metav1.ObjectMeta{Namespace: cluster.Namespace, Name: name}}
	op, err := controllerutil.CreateOrUpdate(context.TODO(), r.client, file, func(e runtime.Object) error {
		f := e.(*resourcesv1alpha1.File)
		if f.CreationTimestamp.IsZero() {
			_ = controllerutil.SetControllerReference(cluster, f, r.scheme)
		}
		f.Spec.Host = cluster.Spec.Controllers[0].Name
		f.Spec.Path = "/root/kubeadm-config.yaml"
		f.Spec.Content = newConfig
		return nil
	})
	if err != nil {
		return pkgerrors.Wrapf(err, "could not reconcile file %s/%s", cluster.Namespace, name)
	}
	if op != controllerutil.OperationResultNone {
		log.Info(fmt.Sprintf("reconciled File %s/%s operation %s", cluster.Namespace, name, op))
	}

	return nil
}

func (r *ReconcileCluster) ensureAdminConf(cluster *kubeadmv1alpha1.Cluster, update *bool) error {
	name := cluster.Name + "-adminconf"
	fc := &resourcesv1alpha1.FileContent{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: cluster.Namespace, Name: name}, fc)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if err != nil {
		fc := &resourcesv1alpha1.FileContent{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: cluster.Namespace,
				Name:      name,
			},
			Spec: resourcesv1alpha1.FileContentSpec{
				Host: cluster.Spec.Controllers[0].Name,
				Path: "/etc/kubernetes/admin.conf",
				Requires: &resourcesv1alpha1.Requires{{
					Command: &resourcesv1alpha1.RequireCommand{
						Name: nameForKubeadmInitCommand(cluster)},
				}},
			},
		}
		_ = controllerutil.SetControllerReference(cluster, fc, r.scheme)
		if err := r.client.Create(context.TODO(), fc); err != nil {
			return err
		}
	} else {
		if fc.Status.Done {
			if cluster.Status.AdminConf != fc.Status.Content {
				cluster.Status.AdminConf = fc.Status.Content
				*update = true
			}
		}
	}

	return nil
}

func nameForKubeadmInitCommand(cluster *kubeadmv1alpha1.Cluster) string {
	return cluster.Name + "-kubeadm-" + cluster.Spec.Controllers[0].Name
}

func (r *ReconcileCluster) runKubeadmOnFirstController(cluster *kubeadmv1alpha1.Cluster, update *bool) error {
	name := nameForKubeadmInitCommand(cluster)
	cmd := &resourcesv1alpha1.FileContent{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: cluster.Namespace, Name: name}, cmd)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	if err != nil {
		fc := &resourcesv1alpha1.Command{
			ObjectMeta: metav1.ObjectMeta{
				Namespace: cluster.Namespace,
				Name:      name,
			},
			Spec: resourcesv1alpha1.CommandSpec{
				Host:    cluster.Spec.Controllers[0].Name,
				Command: "kubeadm init --config=/root/kubeadm-config.yaml --ignore-preflight-errors=all",
				Requires: &resourcesv1alpha1.Requires{{
					File: &resourcesv1alpha1.RequireFile{
						Name: nameForKubeadmConfigFile(cluster)},
				}},
			},
		}
		_ = controllerutil.SetControllerReference(cluster, fc, r.scheme)
		if err := r.client.Create(context.TODO(), fc); err != nil && !errors.IsAlreadyExists(err) { // TODO: ?!
			return err
		}

		return nil
	}

	return nil
}

func (r *ReconcileCluster) ensurePKIStatus(cluster *kubeadmv1alpha1.Cluster, update *bool) error {
	pkifiles := map[string]string{
		"cacert":         "/etc/kubernetes/pki/ca.crt",
		"cakey":          "/etc/kubernetes/pki/ca.key",
		"sakey":          "/etc/kubernetes/pki/sa.key",
		"sapub":          "/etc/kubernetes/pki/sa.pub",
		"frontproxykey":  "/etc/kubernetes/pki/front-proxy-ca.key",
		"frontproxycert": "/etc/kubernetes/pki/front-proxy-ca.crt",
		"etcdcert":       "/etc/kubernetes/pki/etcd/ca.crt",
		"etcdkey":        "/etc/kubernetes/pki/etcd/ca.key",
	}

	pkifields := map[string]*string{
		"cacert":         &cluster.Status.PKI.CACert,
		"cakey":          &cluster.Status.PKI.CAKey,
		"sakey":          &cluster.Status.PKI.SAPrivate,
		"sapub":          &cluster.Status.PKI.SAPublic,
		"frontproxykey":  &cluster.Status.PKI.FrontProxyKey,
		"frontproxycert": &cluster.Status.PKI.FrontProxyCert,
		"etcdcert":       &cluster.Status.PKI.EtcdCert,
		"etcdkey":        &cluster.Status.PKI.EtcdKey,
	}

	for k, _ := range pkifields {
		fcname := fmt.Sprintf("%s-pki-%s", cluster.Name, k)
		fc := &resourcesv1alpha1.FileContent{}
		err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: cluster.Namespace, Name: fcname}, fc)
		if err != nil && !errors.IsNotFound(err) {
			return err
		}
		if err != nil {
			fc := &resourcesv1alpha1.FileContent{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: cluster.Namespace,
					Name:      fcname,
				},
				Spec: resourcesv1alpha1.FileContentSpec{
					Host: cluster.Spec.Controllers[0].Name,
					Path: pkifiles[k],
					Requires: &resourcesv1alpha1.Requires{{
						Command: &resourcesv1alpha1.RequireCommand{
							Name: nameForKubeadmInitCommand(cluster)},
					}},
				},
			}
			_ = controllerutil.SetControllerReference(cluster, fc, r.scheme) // cannot fail

			if err := r.client.Create(context.TODO(), fc); err != nil {
				return err
			}
		} else {
			if fc.Status.Done {
				if *pkifields[k] != fc.Status.Content {
					*pkifields[k] = fc.Status.Content
					*update = true
				}
			}
		}
	}

	return nil
}

func ParseKubeadmOutput(in string) (string, string) {
	re := regexp.MustCompile("--token (\\S+)\\s*--discovery-token-ca-cert-hash\\s*(\\S+)\\s*$")
	ar := re.FindStringSubmatch(in)
	if len(ar) == 3 {
		return ar[1], ar[2]
	}
	return "", ""
}

func (r *ReconcileCluster) ensureTokenStatus(cluster *kubeadmv1alpha1.Cluster, update *bool) error {
	cmd := &resourcesv1alpha1.Command{}
	err := r.client.Get(context.TODO(), types.NamespacedName{Namespace: cluster.Namespace, Name: nameForKubeadmInitCommand(cluster)}, cmd)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		} else {
			return pkgerrors.Wrap(err, "could not get command")
		}
	}

	if !cmd.Status.Done {
		return nil
	}

	token, hash := ParseKubeadmOutput(cmd.Status.Stdout)
	if token == "" || hash == "" {
		return nil // not ready
	}

	cluster.Status.JoinToken = token
	cluster.Status.DiscoveryTokenCACertHash = hash
	*update = true

	return nil
}

func nameForNetworking(cluster *kubeadmv1alpha1.Cluster) string {
	return cluster.Name + "-networking"
}

func (r *ReconcileCluster) ensureNetworking(cluster *kubeadmv1alpha1.Cluster, update *bool) error {
	name := nameForNetworking(cluster)
	cmd := &resourcesv1alpha1.Command{ObjectMeta: metav1.ObjectMeta{Namespace: cluster.Namespace, Name: name}}
	op, err := controllerutil.CreateOrUpdate(context.TODO(), r.client, cmd, func(e runtime.Object) error {
		cmd := e.(*resourcesv1alpha1.Command)
		if cmd.CreationTimestamp.IsZero() {
			_ = controllerutil.SetControllerReference(cluster, cmd, r.scheme)
		}
		cmd.Spec.Host = cluster.Spec.Controllers[0].Name
		cmd.Spec.Command = `KUBECONFIG=/etc/kubernetes/admin.conf kubectl apply -f https://cloud.weave.works/k8s/v1.10/net.yaml`
		cmd.Spec.Requires = &resourcesv1alpha1.Requires{{
			Command: &resourcesv1alpha1.RequireCommand{
				Name: nameForKubeadmInitCommand(cluster)}}}
		return nil
	})
	if err != nil {
		return pkgerrors.Wrapf(err, "could not reconcile file %s/%s", cluster.Namespace, name)
	}
	if op != controllerutil.OperationResultNone {
		log.Info(fmt.Sprintf("reconciled File %s/%s operation %s", cluster.Namespace, name, op))
	}
	return nil
}
