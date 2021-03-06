package kubeconfig

import (
	"fmt"
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var _ = Describe("Pkg/Kubeconfig/Read", func() {
	It("Should fail if kubeconfig doesn't exist", func() {
		config, err := Read("/some/nonexistent/path")

		Expect(config).To(BeNil())
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("Pkg/Kubeconfig/List", func() {
	It("Should return all context names", func() {
		contexts := []string{}
		k := mockConfig(3)
		for context := range k.Contexts {
			contexts = append(contexts, context)
		}
		sort.Strings(contexts)
		Expect(contexts).To(Equal([]string{"test", "test-1", "test-2"}))
	})
})

var _ = Describe("Pkg/Kubeconfig/Export", func() {
	It("Should return a single usable config when given a context name", func() {
		contextName := "test-3"
		config := clientcmdapi.NewConfig()
		config.Clusters[contextName] = &clientcmdapi.Cluster{
			LocationOfOrigin:         "/home/user/.kube/config",
			Server:                   fmt.Sprintf("https://example-%s.com:6443", contextName),
			InsecureSkipTLSVerify:    true,
			CertificateAuthority:     "bbbbbbbbbbbb",
			CertificateAuthorityData: []byte("bbbbbbbbbbbb"),
		}
		config.AuthInfos[contextName] = &clientcmdapi.AuthInfo{
			LocationOfOrigin: "/home/user/.kube/config",
			Token:            fmt.Sprintf("bbbbbbbbbbbb-%s", contextName),
		}
		config.Contexts[contextName] = &clientcmdapi.Context{
			LocationOfOrigin: "/home/user/.kube/config",
			Cluster:          contextName,
			AuthInfo:         contextName,
			Namespace:        "default",
		}
		config.CurrentContext = contextName

		k := mockConfig(5)

		// extract the one config out of the mocked configs
		result, err := k.Export(contextName)

		Expect(result).To(Equal(config))
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should fail if context doesn't exist", func() {
		contextName := "test-7"
		k := mockConfig(5)

		result, err := k.Export(contextName)

		Expect(result).To(BeNil())
		Expect(err).Should(HaveOccurred())
	})
})

var _ = Describe("Pkg/Kubeconfig/GetContent", func() {
	It("Should properly convert a config into bytes that can be written and used as a separate kubeconfig", func() {
		contextName := "test-3"
		k := mockConfig(5)

		// extract the bytes content
		content, err := k.GetContent(contextName)

		Expect(content).ToNot(BeEmpty()) // this helps with the validity check below so we don't get a valid empty config
		Expect(err).ShouldNot(HaveOccurred())

		// convert back to a config as a validity check
		config, err := clientcmd.Load(content)

		Expect(config).ToNot(BeNil())
		Expect(err).ShouldNot(HaveOccurred())
	})
})
