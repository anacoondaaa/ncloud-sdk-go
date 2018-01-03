package sdk_test

import (
	. "github.com/NaverCloudPlatform/ncloud-sdk-go/sdk"
	gock "gopkg.in/h2non/gock.v1"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Public IP Instance List", func() {
	Describe("Get all Public IP Instance List", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Get("/server").
				Reply(200).BodyString(`<getPublicIpInstanceListResponse>
						<requestId>9503bfc0-1802-42b5-9a56-bc640897b6de</requestId>
						<returnCode>0</returnCode>
						<returnMessage>success</returnMessage>
						<totalRows>1</totalRows>
						<publicIpInstanceList>
							<publicIpInstance>
								<publicIpInstanceNo>565346</publicIpInstanceNo>
								<publicIp>210.89.181.250</publicIp>
								<publicIpDescription></publicIpDescription>
								<createDate>2017-12-18T15:16:33+0900</createDate>
								<publicIpInstanceStatusName>used</publicIpInstanceStatusName>
								<publicIpInstanceStatus>
									<code>USED</code>
									<codeName>NET USED state</codeName>
								</publicIpInstanceStatus>
								<publicIpInstanceOperation>
									<code>NULL</code>
									<codeName>NET NULL OP</codeName>
								</publicIpInstanceOperation>
								<publicIpKindType>
									<code>GEN</code>
									<codeName>General</codeName>
								</publicIpKindType>
								<serverInstanceAssociatedWithPublicIp>
									<serverInstanceNo>565284</serverInstanceNo>
									<serverName>svr-9bb8d497d29584a</serverName>
									<serverDescription></serverDescription>
									<cpuCount>2</cpuCount>
									<memorySize>4294967296</memorySize>
									<baseBlockStorageSize>53687091200</baseBlockStorageSize>
									<platformType>
										<code>LNX64</code>
										<codeName>Linux 64 Bit</codeName>
									</platformType>
									<loginKeyName>packer-1513575006</loginKeyName>
									<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
									<publicIp>210.89.181.250</publicIp>
									<privateIp>10.33.2.206</privateIp>
									<serverImageName>centos-5.11-64</serverImageName>
									<serverInstanceStatus>
										<code>RUN</code>
										<codeName>Server run state</codeName>
									</serverInstanceStatus>
									<serverInstanceOperation>
										<code>NULL</code>
										<codeName>Server NULL OP</codeName>
									</serverInstanceOperation>
									<serverInstanceStatusName>running</serverInstanceStatusName>
									<createDate>2017-12-18T14:30:08+0900</createDate>
									<uptime>2017-12-18T14:35:09+0900</uptime>
									<serverImageProductCode>SPSW0LINUX000043</serverImageProductCode>
									<serverProductCode>SPSVRSTAND000004</serverProductCode>
									<isProtectServerTermination>false</isProtectServerTermination>
									<portForwardingPublicIp>211.249.60.222</portForwardingPublicIp>
									<zone>
										<zoneNo>2</zoneNo>
										<zoneName>KR-1</zoneName>
										<zoneDescription>가산 NANG zone</zoneDescription>
									</zone>
									<region>
										<regionNo>1</regionNo>
										<regionCode>KR</regionCode>
										<regionName>Korea</regionName>
									</region>
									<baseBlockStorageDiskType>
										<code>NET</code>
										<codeName>Network Storage</codeName>
									</baseBlockStorageDiskType>
									<baseBlockStroageDiskDetailType>
										<code>HDD</code>
										<codeName>HDD</codeName>
									</baseBlockStroageDiskDetailType>
									<userData></userData>
									<accessControlGroupList>
										<accessControlGroup>
											<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
											<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
											<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
											<isDefault>true</isDefault>
											<createDate>2017-02-23T10:25:39+0900</createDate>
										</accessControlGroup>
									</accessControlGroupList>
								</serverInstanceAssociatedWithPublicIp>
							</publicIpInstance>
						</publicIpInstanceList>
					</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get Public IP Instance List", func() {
			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(nil)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("9503bfc0-1802-42b5-9a56-bc640897b6de"))
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))

			publicIPInstance := result.PublicIPInstanceList[0]
			Expect(publicIPInstance.PublicIPInstanceNo).To(Equal("565346"))
			Expect(publicIPInstance.PublicIP).To(Equal("210.89.181.250"))
			Expect(publicIPInstance.CreateDate).To(Equal("2017-12-18T15:16:33+0900"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("used"))
			Expect(publicIPInstance.PublicIPInstanceStatus.Code).To(Equal("USED"))
			Expect(publicIPInstance.PublicIPInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(publicIPInstance.PublicIPInstanceOperation.Code).To(Equal("NULL"))
			Expect(publicIPInstance.PublicIPInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(publicIPInstance.PublicIPKindType.Code).To(Equal("GEN"))
			Expect(publicIPInstance.PublicIPKindType.CodeName).To(Equal("General"))
			Expect(publicIPInstance.ServerInstance.ServerInstanceNo).To(Equal("565284"))
		})
	})

	Describe("Get Public IP Instance List which is parameters", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Get("/server").
				Reply(200).BodyString(`<getPublicIpInstanceListResponse>
					<requestId>3c0668ae-6c6d-4b60-9bb5-874204f81802</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>1</totalRows>
					<publicIpInstanceList>
						<publicIpInstance>
							<publicIpInstanceNo>565346</publicIpInstanceNo>
							<publicIp>210.89.181.250</publicIp>
							<publicIpDescription></publicIpDescription>
							<createDate>2017-12-18T15:16:33+0900</createDate>
							<publicIpInstanceStatusName>used</publicIpInstanceStatusName>
							<publicIpInstanceStatus>
								<code>USED</code>
								<codeName>NET USED state</codeName>
							</publicIpInstanceStatus>
							<publicIpInstanceOperation>
								<code>NULL</code>
								<codeName>NET NULL OP</codeName>
							</publicIpInstanceOperation>
							<publicIpKindType>
								<code>GEN</code>
								<codeName>General</codeName>
							</publicIpKindType>
							<serverInstanceAssociatedWithPublicIp>
								<serverInstanceNo>565284</serverInstanceNo>
								<serverName>svr-9bb8d497d29584a</serverName>
								<serverDescription></serverDescription>
								<cpuCount>2</cpuCount>
								<memorySize>4294967296</memorySize>
								<baseBlockStorageSize>53687091200</baseBlockStorageSize>
								<platformType>
									<code>LNX64</code>
									<codeName>Linux 64 Bit</codeName>
								</platformType>
								<loginKeyName>packer-1513575006</loginKeyName>
								<isFeeChargingMonitoring>false</isFeeChargingMonitoring>
								<publicIp>210.89.181.250</publicIp>
								<privateIp>10.33.2.206</privateIp>
								<serverImageName>centos-5.11-64</serverImageName>
								<serverInstanceStatus>
									<code>RUN</code>
									<codeName>Server run state</codeName>
								</serverInstanceStatus>
								<serverInstanceOperation>
									<code>NULL</code>
									<codeName>Server NULL OP</codeName>
								</serverInstanceOperation>
								<serverInstanceStatusName>running</serverInstanceStatusName>
								<createDate>2017-12-18T14:30:08+0900</createDate>
								<uptime>2017-12-18T14:35:09+0900</uptime>
								<serverImageProductCode>SPSW0LINUX000043</serverImageProductCode>
								<serverProductCode>SPSVRSTAND000004</serverProductCode>
								<isProtectServerTermination>false</isProtectServerTermination>
								<portForwardingPublicIp>211.249.60.222</portForwardingPublicIp>
								<zone>
									<zoneNo>2</zoneNo>
									<zoneName>KR-1</zoneName>
									<zoneDescription>가산 NANG zone</zoneDescription>
								</zone>
								<region>
									<regionNo>1</regionNo>
									<regionCode>KR</regionCode>
									<regionName>Korea</regionName>
								</region>
								<baseBlockStorageDiskType>
									<code>NET</code>
									<codeName>Network Storage</codeName>
								</baseBlockStorageDiskType>
								<baseBlockStroageDiskDetailType>
									<code>HDD</code>
									<codeName>HDD</codeName>
								</baseBlockStroageDiskDetailType>
								<userData></userData>
								<accessControlGroupList>
									<accessControlGroup>
										<accessControlGroupConfigurationNo>4964</accessControlGroupConfigurationNo>
										<accessControlGroupName>ncloud-default-acg</accessControlGroupName>
										<accessControlGroupDescription>Default AccessControlGroup</accessControlGroupDescription>
										<isDefault>true</isDefault>
										<createDate>2017-02-23T10:25:39+0900</createDate>
									</accessControlGroup>
								</accessControlGroupList>
							</serverInstanceAssociatedWithPublicIp>
						</publicIpInstance>
					</publicIpInstanceList>
				</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get One Public IP Instance List", func() {
			reqParams := new(RequestPublicIPInstanceList)
			reqParams.PublicIPInstanceNoList = []string{"565346"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("3c0668ae-6c6d-4b60-9bb5-874204f81802"))
			Expect(result.TotalRows).To(Equal(1))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(1))

			publicIPInstance := result.PublicIPInstanceList[0]
			Expect(publicIPInstance.PublicIPInstanceNo).To(Equal("565346"))
			Expect(publicIPInstance.PublicIP).To(Equal("210.89.181.250"))
			Expect(publicIPInstance.CreateDate).To(Equal("2017-12-18T15:16:33+0900"))
			Expect(publicIPInstance.PublicIPInstanceStatusName).To(Equal("used"))
			Expect(publicIPInstance.PublicIPInstanceStatus.Code).To(Equal("USED"))
			Expect(publicIPInstance.PublicIPInstanceStatus.CodeName).To(Equal("NET USED state"))
			Expect(publicIPInstance.PublicIPInstanceOperation.Code).To(Equal("NULL"))
			Expect(publicIPInstance.PublicIPInstanceOperation.CodeName).To(Equal("NET NULL OP"))
			Expect(publicIPInstance.PublicIPKindType.Code).To(Equal("GEN"))
			Expect(publicIPInstance.PublicIPKindType.CodeName).To(Equal("General"))
			Expect(publicIPInstance.ServerInstance.ServerInstanceNo).To(Equal("565284"))
		})
	})

	Describe("There is no Public IP Instance list", func() {
		BeforeEach(func() {
			gock.New("https://api.ncloud.com").
				Get("/server").
				Reply(200).BodyString(`<getPublicIpInstanceListResponse>
					<requestId>f7791c1d-941f-4968-bc24-a7224e16d39f</requestId>
					<returnCode>0</returnCode>
					<returnMessage>success</returnMessage>
					<totalRows>0</totalRows>
					<publicIpInstanceList/>
				</getPublicIpInstanceListResponse>`)
		})
		AfterEach(func() {
			gock.Off()
		})
		It("should get No Public IP Instance list", func() {
			reqParams := new(RequestPublicIPInstanceList)
			reqParams.PublicIPInstanceNoList = []string{"5635346"}

			conn := NewConnection(accessKey, secretKey)
			result, err := conn.GetPublicIPInstanceList(reqParams)

			Expect(err).To(BeNil())
			Expect(result.RequestID).To(Equal("f7791c1d-941f-4968-bc24-a7224e16d39f"))
			Expect(result.TotalRows).To(Equal(0))
			Expect(result.ReturnCode).To(Equal(0))
			Expect(result.ReturnMessage).To(Equal("success"))
			Expect(len(result.PublicIPInstanceList)).To(Equal(0))
		})
	})
})
