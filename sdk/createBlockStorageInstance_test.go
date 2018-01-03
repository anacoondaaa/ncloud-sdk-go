package sdk_test

import (
	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Create Block Storage Instance", func() {
	Describe("Create Block Storage Instance", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Get("/server").
				Reply(200).BodyString(`<createBlockStorageInstanceResponse>
					<requestId>4920c659-d272-4357-b632-fc3a25fa80b2</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<blockStorageInstanceList>
						<blockStorageInstance>
							<blockStorageInstanceNo>558934</blockStorageInstanceNo>
							<serverInstanceNo>558779</serverInstanceNo>
							<serverName>svr-9bb68126c6a0ce6</serverName>
							<blockStorageType>
								<code>SVRBS</code>
								<codeName>Server BS</codeName>
							</blockStorageType>
							<blockStorageName>blockstorage</blockStorageName>
							<blockStorageSize>53687091200</blockStorageSize>
							<deviceName></deviceName>
							<blockStorageProductCode>SPBSTBSTAD000002</blockStorageProductCode>
							<blockStorageInstanceStatus>
								<code>INIT</code>
								<codeName>Block storage INIT state</codeName>
							</blockStorageInstanceStatus>
							<blockStorageInstanceOperation>
								<code>NULL</code>
								<codeName>Block Storage NULLOP</codeName>
							</blockStorageInstanceOperation>
							<blockStorageInstanceStatusName>initialized</blockStorageInstanceStatusName>
							<createDate>2017-12-08T18:25:07+0900</createDate>
							<blockStorageInstanceDescription>Block Storage Instance by SDK go</blockStorageInstanceDescription>
							<diskType>
								<code>NET</code>
								<codeName>Network Storage</codeName>
							</diskType>
							<diskDetailType>
								<code>HDD</code>
								<codeName>HDD</codeName>
							</diskDetailType>
						</blockStorageInstance>
					</blockStorageInstanceList>
				</createBlockStorageInstanceResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should create Block Storage Instance", func() {
			reqParams := new(RequestBlockStorageInstance)
			reqParams.BlockStorageName = "blockstorage"
			reqParams.BlockStorageSize = 50
			reqParams.BlockStorageDescription = "Block Storage Instance by SDK go"
			reqParams.ServerInstanceNo = "558779"

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.CreateBlockStorageInstance(reqParams)

			Expect(err).To(BeNil())
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.BlockStorageInstance)).To(Equal(1))

			blockStorageInstance := result.BlockStorageInstance[0]
			Expect(blockStorageInstance.BlockStorageInstanceNo).To(Equal("558934"))
			Expect(blockStorageInstance.ServerInstanceNo).To(Equal("558779"))
			Expect(blockStorageInstance.ServerName).To(Equal("svr-9bb68126c6a0ce6"))
			Expect(blockStorageInstance.BlockStorageType.Code).To(Equal("SVRBS"))
			Expect(blockStorageInstance.BlockStorageType.CodeName).To(Equal("Server BS"))
			Expect(blockStorageInstance.BlockStorageName).To(Equal("blockstorage"))
			Expect(blockStorageInstance.BlockStorageSize).To(Equal(53687091200))
			Expect(blockStorageInstance.DeviceName).To(Equal(""))
			Expect(blockStorageInstance.BlockStorageProductCode).To(Equal("SPBSTBSTAD000002"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.Code).To(Equal("INIT"))
			Expect(blockStorageInstance.BlockStorageInstanceStatus.CodeName).To(Equal("Block storage INIT state"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.Code).To(Equal("NULL"))
			Expect(blockStorageInstance.BlockStorageInstanceOperation.CodeName).To(Equal("Block Storage NULLOP"))
			Expect(blockStorageInstance.BlockStorageInstanceStatusName).To(Equal("initialized"))
			Expect(blockStorageInstance.CreateDate).To(Equal("2017-12-08T18:25:07+0900"))
			Expect(blockStorageInstance.BlockStorageInstanceDescription).To(Equal("Block Storage Instance by SDK go"))
			Expect(blockStorageInstance.DiskType.Code).To(Equal("NET"))
			Expect(blockStorageInstance.DiskType.CodeName).To(Equal("Network Storage"))
			Expect(blockStorageInstance.DiskDetailType.Code).To(Equal("HDD"))
			Expect(blockStorageInstance.DiskDetailType.CodeName).To(Equal("HDD"))

		})
	})

	Describe("Unable to create Block Storage Instance", func() {
		Context("when invalid ServerInstanceNo is invalid", func() {
			BeforeEach(func() {
				gock.New("https://api.ncloud.com").
					Get("/server").
					Reply(500).BodyString(`<responseError>
						<returnCode>1300</returnCode>
						<returnMessage>Please try your call again later.
					Temporarily out of service.
					If error continue, Please contact our customer service center.</returnMessage>
					</responseError>`)
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail", func() {
				reqParams := new(RequestBlockStorageInstance)
				reqParams.BlockStorageName = "blockstorage"
				reqParams.BlockStorageSize = 50
				reqParams.BlockStorageDescription = "Block Storage Instance by SDK go"
				reqParams.ServerInstanceNo = "558780"

				conn := NewConnection(accessKey, secretKey)
				result, err := conn.CreateBlockStorageInstance(reqParams)

				Expect(err.Error()).To(ContainSubstring("500 Internal Server Error"))
				Expect(result.ReturnCode).To(Equal(1300))
				Expect(result.ReturnMessage).To(Equal(`Please try your call again later.
					Temporarily out of service.
					If error continue, Please contact our customer service center.`))
			})
		})

		Context("when BlockStorageSize is not a multiple of 10 GB", func() {
			It("should fail", func() {
				reqParams := new(RequestBlockStorageInstance)
				reqParams.BlockStorageName = "blockstorage"
				reqParams.BlockStorageSize = 502
				reqParams.BlockStorageDescription = "Block Storage Instance by SDK go"
				reqParams.ServerInstanceNo = "558780"

				conn := NewConnection(accessKey, secretKey)
				_, err := conn.CreateBlockStorageInstance(reqParams)

				Expect(err.Error()).To(ContainSubstring("BlockStorageSize must be a multiple of 10 GB"))
			})
		})

		Context("Duplicate Block Storage Instance name", func() {
			BeforeEach(func() {
				gock.New("https://api.ncloud.com").
					Get("/server").
					Reply(400).BodyString(`<responseError>
						<returnCode>10300</returnCode>
						<returnMessage>Instance name is already in use. please use other name.</returnMessage>
					</responseError>`)
			})
			AfterEach(func() {
				gock.Off()
			})

			It("should fail if (other) user is either operating server or creating Block Storage Instance of the server.", func() {
				reqParams := new(RequestBlockStorageInstance)
				reqParams.BlockStorageName = "blockstorage"
				reqParams.BlockStorageSize = 50
				reqParams.BlockStorageDescription = "Block Storage Instance by SDK go"
				reqParams.ServerInstanceNo = "558780"

				conn := NewConnection(accessKey, secretKey)
				result, err := conn.CreateBlockStorageInstance(reqParams)

				Expect(err.Error()).To(ContainSubstring("400 Bad Request"))
				Expect(result.ReturnCode).To(Equal(10300))
				Expect(result.ReturnMessage).To(Equal("Instance name is already in use. please use other name."))
			})
		})
	})
})
