package controllers

import (
	goctx "context"
	"time"

	admv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	appsv1 "k8s.io/api/apps/v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	sriovnetworkv1 "github.com/openshift/sriov-network-operator/api/v1"
	util "github.com/openshift/sriov-network-operator/test/util"
)

var _ = Describe("Operator", func() {

	BeforeEach(func() {
		// wait for sriov-network-operator to be ready
		config := &sriovnetworkv1.SriovOperatorConfig{}
		err := util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
		Expect(err).NotTo(HaveOccurred())

		*config.Spec.EnableOperatorWebhook = true
		*config.Spec.EnableInjector = true
		config.Spec.ConfigDaemonNodeSelector = nil

		err = k8sClient.Update(goctx.TODO(), config)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("When is up", func() {
		It("should have a default operator config", func() {
			config := &sriovnetworkv1.SriovOperatorConfig{}
			err := util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			Expect(*config.Spec.EnableOperatorWebhook).To(Equal(true))
			Expect(*config.Spec.EnableInjector).To(Equal(true))
			Expect(config.Spec.ConfigDaemonNodeSelector).Should(BeNil())
		})

		It("should have webhook enable", func() {
			mutateCfg := &admv1beta1.MutatingWebhookConfiguration{}
			err := util.WaitForNamespacedObject(mutateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			validateCfg := &admv1beta1.ValidatingWebhookConfiguration{}
			err = util.WaitForNamespacedObject(validateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())
		})

		DescribeTable("should have daemonset enabled by default",
			func(dsName string) {
				// wait for sriov-network-operator to be ready
				daemonSet := &appsv1.DaemonSet{}
				err := util.WaitForNamespacedObject(daemonSet, k8sClient, testNamespace, dsName, interval, timeout)
				Expect(err).NotTo(HaveOccurred())
			},
			Entry("operator-webhook", "operator-webhook"),
			Entry("network-resources-injector", "network-resources-injector"),
			Entry("sriov-network-config-daemon", "sriov-network-config-daemon"),
		)

		It("should be able to turn network-resources-injector on/off", func() {
			By("set disable to enableInjector")
			config := &sriovnetworkv1.SriovOperatorConfig{}
			err := util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			*config.Spec.EnableInjector = false
			err = k8sClient.Update(goctx.TODO(), config)
			Expect(err).NotTo(HaveOccurred())

			daemonSet := &appsv1.DaemonSet{}
			err = util.WaitForNamespacedObjectDeleted(daemonSet, k8sClient, testNamespace, "network-resources-injector", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			mutateCfg := &admv1beta1.MutatingWebhookConfiguration{}
			err = util.WaitForNamespacedObjectDeleted(mutateCfg, k8sClient, testNamespace, "network-resources-injector-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			By("set enable to enableInjector")
			err = util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			*config.Spec.EnableInjector = true
			err = k8sClient.Update(goctx.TODO(), config)
			Expect(err).NotTo(HaveOccurred())

			daemonSet = &appsv1.DaemonSet{}
			err = util.WaitForNamespacedObject(daemonSet, k8sClient, testNamespace, "network-resources-injector", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			mutateCfg = &admv1beta1.MutatingWebhookConfiguration{}
			err = util.WaitForNamespacedObject(mutateCfg, k8sClient, testNamespace, "network-resources-injector-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())
		})

		It("should be able to turn operator-webhook on/off", func() {

			By("set disable to enableOperatorWebhook")
			config := &sriovnetworkv1.SriovOperatorConfig{}
			err := util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			*config.Spec.EnableOperatorWebhook = false
			err = k8sClient.Update(goctx.TODO(), config)
			Expect(err).NotTo(HaveOccurred())

			daemonSet := &appsv1.DaemonSet{}
			err = util.WaitForNamespacedObjectDeleted(daemonSet, k8sClient, testNamespace, "operator-webhook", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			mutateCfg := &admv1beta1.MutatingWebhookConfiguration{}
			err = util.WaitForNamespacedObjectDeleted(mutateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			validateCfg := &admv1beta1.ValidatingWebhookConfiguration{}
			err = util.WaitForNamespacedObjectDeleted(validateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			By("set disable to enableOperatorWebhook")
			err = util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			*config.Spec.EnableOperatorWebhook = true
			err = k8sClient.Update(goctx.TODO(), config)
			Expect(err).NotTo(HaveOccurred())

			daemonSet = &appsv1.DaemonSet{}
			err = util.WaitForNamespacedObject(daemonSet, k8sClient, testNamespace, "operator-webhook", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			mutateCfg = &admv1beta1.MutatingWebhookConfiguration{}
			err = util.WaitForNamespacedObject(mutateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())

			validateCfg = &admv1beta1.ValidatingWebhookConfiguration{}
			err = util.WaitForNamespacedObject(validateCfg, k8sClient, testNamespace, "operator-webhook-config", interval, timeout)
			Expect(err).NotTo(HaveOccurred())
		})
		PIt("should be able to update the node selector of sriov-network-config-daemon", func() {
			By("specify the configDaemonNodeSelector")
			config := &sriovnetworkv1.SriovOperatorConfig{}
			err := util.WaitForNamespacedObject(config, k8sClient, testNamespace, "default", interval, timeout)
			Expect(err).NotTo(HaveOccurred())
			config.Spec.ConfigDaemonNodeSelector = map[string]string{"node-role.kubernetes.io/worker": ""}
			err = k8sClient.Update(goctx.TODO(), config)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(10 * time.Second)
			daemonSet := &appsv1.DaemonSet{}
			err = util.WaitForDaemonSetReady(daemonSet, k8sClient, namespace, "sriov-network-config-daemon", interval, timeout)
			Expect(err).NotTo(HaveOccurred())
			Expect(daemonSet.Spec.Template.Spec.NodeSelector).To(Equal(config.Spec.ConfigDaemonNodeSelector))
		})
	})
})
