package payload

const (
	MSGIDBASE = "MSGIDBASE"
	// MsgSeperator is used to separate sent messages via NETCONF
	NetconfBegin = "<"
	MsgSeperator = "]]>]]>"

	HelloRsp = `
<hello xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
<capabilities>
<capability>urn:ietf:params:netconf:base:1.1</capability>
<capability>urn:ietf:params:netconf:capability:startup:1.0</capability>
<capability>urn:ietf:params:o-ran:capability:startup:1.0</capability>
<capability>urn:o-ran:beamforming:1.0</capability>
</capabilities>
<session-id>4</session-id>
</hello>
`

	RpcA3TopoListRsp = `<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
<data>
<topology-list>
<node-array>
<dvcId>1</dvcId>          
	<serial-number>10086</serial-number>
	<mac-address>dr1454rkokd9090ifrf</mac-address>
	<product-model>12358rt</product-model>
	<firmware-version>1.2.3.6</firmware-version>
	<online-state>online</online-state>
	<ip-address>10.2.3.36</ip-address>       
	<north-node>
		<port-number>10</port-number>
		<mac-address>fe4regfre43gr</mac-address>
		<serial-number>10086-10</serial-number>
	</north-node>
</node-array>
<node-array>
<dvcId>2</dvcId>          
	<serial-number>10087</serial-number>
	<mac-address>dvwfew34232</mac-address>
	<product-model>1256230</product-model>
	<firmware-version>2.363.265</firmware-version>
	<online-state>offline</online-state>
	<ip-address>10.5.62.3</ip-address>       
	<north-node>
		<port-number>20</port-number>
		<mac-address>efefewgfw43rtrtf34t</mac-address>
		<serial-number>1008620</serial-number>
	</north-node>
</node-array>
</topology-list>
</data>
</rpc-reply>`

	RpcModulesRsp = `<rpc-reply
xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
<data
	xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
	<modules-state
		xmlns="urn:ietf:params:xml:ns:yang:ietf-yang-library">
		<module>
			<name>o-ran-operations</name>
			<revision>2019-07-03</revision>
			<namespace>urn:o-ran:operations:1.0</namespace>
			<conformance-type>implement</conformance-type>
		</module>
	</modules-state>
</data>
</rpc-reply>`

	RpcCreateSubscriptionReplyMsg = `
<rpc-reply xmlns:netconf="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
<ok/>
</rpc-reply>
`
	RpcReplyMsg = `
<rpc-reply xmlns:netconf="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
<ok/>
</rpc-reply>
`

	DownloadReqRpcReplyMsg = `
	<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
		<status xmlns="urn:o-ran:software-management:1.0">STARTED</status>
		<notification-timeout xmlns="urn:o-ran:software-management:1.0">%d</notification-timeout>
	</rpc-reply>
`

	InstallReqRpcReplyMsg = `
	<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
		<status xmlns="urn:o-ran:software-management:1.0">STARTED</status>
		<notification-timeout xmlns="urn:o-ran:software-management:1.0">30</notification-timeout>
	</rpc-reply>
`
	ActivateReqRpcReplyMsg = `
	<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
		<status xmlns="urn:o-ran:software-management:1.0">STARTED</status>
		<notification-timeout xmlns="urn:o-ran:software-management:1.0">30</notification-timeout>
	</rpc-reply>
`

	RpcReplyMsgGetAlarmList = `
<rpc-reply xmlns="urn:ietf:params:xml:ns:netconf:base:1.0" message-id="MSGIDBASE">
<data>
        <active-alarm-list
            xmlns="urn:o-ran:fm:1.0">
            <active-alarms>
                <fault-id>11</fault-id>
                <fault-source>N71N29N26-RU_RU Local Access Attempt</fault-source>
                <affected-objects>
                    <name>N71N29N26-RU</name>
                </affected-objects>
                <fault-severity>MINOR</fault-severity>
                <is-cleared>false</is-cleared>
                <fault-text>RU Local Access Attempt</fault-text>
                <event-time>2023-02-10T08:48:47+00:00</event-time>
            </active-alarms>
            <active-alarms>
                <fault-id>1</fault-id>
                <fault-source>N71N29N26-RU_High Temp_Minor</fault-source>
                <affected-objects>
                    <name>N71N29N26-RU</name>
                </affected-objects>
                <fault-severity>MINOR</fault-severity>
                <is-cleared>false</is-cleared>
                <fault-text>Node Temperature is high</fault-text>
                <event-time>2023-02-10T08:48:46+00:00</event-time>
            </active-alarms>
        </active-alarm-list>
    </data>
</rpc-reply>
`

	NotifyMsg = `
<notification xmlns:netconf="urn:ietf:params:xml:ns:netconf:base:1.0">
<ok/>
<timestamp>%s</timestamp>
<duid>%d</duid>
<notifId>%d</notifId>
</notification>
`

	DownNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<eventTime>2019-09-25T02:35:21Z</eventTime>
<download-event xmlns="urn:o-ran:software-management:1.0">
<file-name>Mavenir_RD4401_1909051644_Cut1.ffw</file-name>
<status>COMPLETED</status>
</download-event>
</notification>
`

	InstalNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<eventTime>2019-09-25T02:35:21Z</eventTime>
<install-event xmlns="urn:o-ran:software-management:1.0">
<slot-name>passiveswpkg</slot-name>
<status>COMPLETED</status>
</install-event>
</notification>
`

	ActivateNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<eventTime>2019-09-25T02:35:21Z</eventTime>
<activation-event xmlns="urn:o-ran:software-management:1.0">
<slot-name>passiveswpkg</slot-name>
<status>COMPLETED</status>
<return-code>0</return-code>
</activation-event>
</notification>
`

	AlaramNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<alarm-notif xmlns="urn:o-ran:software-management:1.0">
<fault-id>1200</fault-id>
<FaultSource>alarmtest</FaultSource>
<eventTime>2019-09-25T02:35:21Z</eventTime>
</alarm-notif>
</notification>
`
	A3TopoEventNtf = `<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<topology-list-changed>
<test>1</test>
</topology-list-changed>
</notification>`

	OperTrxStateNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<oper_trx_state  xmlns="urn:mavenir:npn-state-notify:1.0">      
<nodeState>SUCCESS/Fail </nodeState>
<serialNumber>m072810</serialNumber>
<operation>transactionAction </operation>
<transactionID>20220119063423232</transactionID>
<msgInfo></msgInfo>
<timestamp>1627547100314321258</timestamp>
</oper_trx_state>
</notification>
`

	OperNodeStateNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<oper_node_state xmlns="urn:mavenir:npn-state-notify:1.0">      
	<type>du</type>
	<nodeName>du_host_1</nodeName>
	<nodeIP>10.2.60.175</nodeIP>
	<nodeState>SUCCESS/Fail </nodeState>
               <serialNumber>m072810</serialNumber>
               <operation>AddNode/DelNode</operation>
               <transactionID>20220119063423232</transactionID>
               <msgInfo></msgInfo>
               <timestamp>1627547100314321258</timestamp>
    </oper_node_state>

</notification>
`

	MonitorNodeStateNtf = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
<monitor_node_state  xmlns="urn:mavenir:npn-state-notify:1.0">      
	<name>5gc_host_1</name>
	<nodeIP>10.2.60.175</nodeIP>
	<roles>master/worker</roles>
	<nodeState>Ready/Not Ready </nodeState>
	<serialNumber>m072810</serialNumber>
	<msgInfo></msgInfo>
	<timestamp>1627547100314321258</timestamp>
</monitor_node_state>
</notification>
`
)

const (
	MeasurementResultStats = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
		<measurement-result-stats xmlns="urn:o-ran:performance-management:1.0">
		<rx-window-stats>
            <measurement-object>RX_ON_TIME</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>19319663</count>
        </rx-window-stats>
        <rx-window-stats>
            <measurement-object>RX_EARLY</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>0</count>
        </rx-window-stats>
        <rx-window-stats>
            <measurement-object>RX_LATE</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>329</count>
        </rx-window-stats>
        <rx-window-stats>
            <measurement-object>RX_CORRUPT</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>0</count>
        </rx-window-stats>
        <rx-window-stats>
            <measurement-object>RX_DUPL</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>0</count>
        </rx-window-stats>
        <rx-window-stats>
            <measurement-object>RX_TOTAL</measurement-object>
            <start-time>2021-03-29T06:05:00+00:00</start-time>
            <end-time>2021-03-29T06:10:00+00:00</end-time>
            <name>N71N29N26-RU</name>
            <count>19319992</count>
        </rx-window-stats>
		<transceiver-stats>
		  <measurement-object>RX_POWER</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <transceiver-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>0.5194</value>
			  <time>2020-07-10T06:09:05+01:00</time>
			</min>
			<max>
			  <value>0.523</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>0.523</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>0.5194</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </transceiver-measurement-result>
		</transceiver-stats>
		<transceiver-stats>
		  <measurement-object>TX_POPWER</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <transceiver-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>0.5879</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>0.5927</value>
			  <time>2020-07-10T06:08:50+01:00</time>
			</max>
			<first>
			  <value>0.5879</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>0.5879</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </transceiver-measurement-result>
		</transceiver-stats>
		<vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
		  <measurement-object>RSSI</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <vendor-specific-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		  <vendor-specific-measurement-result>
			<object-unit-id>1</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		  <vendor-specific-measurement-result>
			<object-unit-id>16</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		  <vendor-specific-measurement-result>
			<object-unit-id>17</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		</vendor-specific-stats>
		<vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
		  <measurement-object>VSWR</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <vendor-specific-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>1.0</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>1.0</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>1.0</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>1.0</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		</vendor-specific-stats>
		<vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
		  <measurement-object>TX_OUTPOWER</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <vendor-specific-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		</vendor-specific-stats>
		<vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
		  <measurement-object>TX_INPOWER</measurement-object>
		  <start-time>2020-07-10T06:08:40+01:00</start-time>
		  <end-time>2020-07-10T06:10:00+01:00</end-time>
		  <vendor-specific-measurement-result>
			<object-unit-id>0</object-unit-id>
			<min>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</min>
			<max>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</max>
			<first>
			  <value>-999.9</value>
			  <time>2020-07-10T06:08:40+01:00</time>
			</first>
			<latest>
			  <value>-999.9</value>
			  <time>2020-07-10T06:10:00+01:00</time>
			</latest>
		  </vendor-specific-measurement-result>
		</vendor-specific-stats>
	  </measurement-result-stats>	 
</notification>
		`

	MeasurementResultStatsWithTimeStamp = `
<notification xmlns="urn:ietf:params:xml:ns:netconf:notification:1.0">
    <measurement-result-stats xmlns="urn:o-ran:performance-management:1.0">
    <rx-window-stats>
    <measurement-object>RX_TOTAL</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <count>28802600</count>
    </rx-window-stats>
    <rx-window-stats>
    <measurement-object>RX_ON_TIME</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <count>99</count>
    </rx-window-stats>
    <rx-window-stats>
    <measurement-object>RX_LATE</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <count>77</count>
    </rx-window-stats>
    <rx-window-stats>
    <measurement-object>RX_EARLY</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <count>28801920</count>
    </rx-window-stats>
    <transceiver-stats>
    <measurement-object>RX_POWER</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <transceiver-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>0.333</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>0.555</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>0.5</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>0.6</value>
    <time>${reporttime}</time>
    </latest>
    </transceiver-measurement-result>
    </transceiver-stats>
    <transceiver-stats>
    <measurement-object>TX_POWER</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <transceiver-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>0.5879</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>0.5927</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>0.5879</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>0.5879</value>
    <time>${reporttime}</time>
    </latest>
    </transceiver-measurement-result>
    </transceiver-stats>
    <transceiver-stats>
    <measurement-object>TX_BIAS_COUNT</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <transceiver-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>6.378</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>6.414</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>6.414</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>6.414</value>
    <time>${reporttime}</time>
    </latest>
    </transceiver-measurement-result>
    </transceiver-stats>
    <transceiver-stats>
    <measurement-object>VOLTAGE</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <transceiver-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>3321.5999</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>3324.0</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>3324.0</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>3324.0</value>
    <time>${reporttime}</time>
    </latest>
    </transceiver-measurement-result>
    </transceiver-stats>
    <transceiver-stats>
    <measurement-object>TEMPERATURE</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <transceiver-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>39.375</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>39.668</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>39.375</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>39.5</value>
    <time>${reporttime}</time>
    </latest>
    </transceiver-measurement-result>
    </transceiver-stats>
    <vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
    <measurement-object>RSSI</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <vendor-specific-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>1</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>16</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>17</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    </vendor-specific-stats>
    <vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
    <measurement-object>VSWR</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <vendor-specific-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>1.0</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>1.0</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>1.0</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>1.0</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>1</object-unit-id>
    <min>
    <value>1.0</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>1.0</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>1.0</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>1.0</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    </vendor-specific-stats>
    <vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
    <measurement-object>TX_OUTPOWER</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <vendor-specific-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>1</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    </vendor-specific-stats>
    <vendor-specific-stats xmlns="urn:vs:performance-management:1.0">
    <measurement-object>TX_INPOWER</measurement-object>
    <start-time>${reporttime}</start-time>
    <end-time>${reporttime}</end-time>
    <vendor-specific-measurement-result>
    <object-unit-id>0</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    <vendor-specific-measurement-result>
    <object-unit-id>1</object-unit-id>
    <min>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </min>
    <max>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </max>
    <first>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </first>
    <latest>
    <value>-999.9</value>
    <time>${reporttime}</time>
    </latest>
    </vendor-specific-measurement-result>
    </vendor-specific-stats>
    </measurement-result-stats>
    </notification>]]>]]>`

	RPCCReplyMsgGetCfg = `<rpc-reply message-id="MSGIDBASE"

  xmlns:nc="urn:ietf:params:xml:ns:netconf:base:1.0"
  last-modified="2022-04-26T12:19:38Z"
  xmlns="urn:ietf:params:xml:ns:netconf:base:1.0">
  <data>
    <software-inventory xmlns="urn:o-ran:software-management:1.0">
      <software-slot>
        <name>Default</name>
        <status>VALID</status>
        <active>true</active>
        <running>true</running>
        <access>READ_ONLY</access>
        <product-code>G18RRH-43-01B</product-code>
        <build-id>1</build-id>
        <build-name>image1</build-name>
        <build-version>v1.00</build-version>
        <files>
          <name>test1_file1</name>
          <version>v1.00</version>
          <local-path>/tmp/file1</local-path>
        </files>
        <files>
          <name/>
          <version/>
          <local-path/>
        </files>
      </software-slot>
      <software-slot>
        <name>slot2</name>
        <status>VALID</status>
        <active>false</active>
        <running>false</running>
        <access>READ_WRITE</access>
        <product-code/>
        <build-id/>
        <build-name/>
        <build-version/>
        <files>
          <name/>
          <version/>
          <local-path/>
        </files>
        <files>
          <name/>
          <version/>
          <local-path/>
        </files>
      </software-slot>
      <software-slot>
        <name>slot3</name>
        <status>VALID</status>
        <active>false</active>
        <running>false</running>
        <access>READ_WRITE</access>
        <product-code/>
        <build-id/>
        <build-name/>
        <build-version/>
        <files>
          <name/>
          <version/>
          <local-path/>
        </files>
        <files>
          <name/>
          <version/>
          <local-path/>
        </files>
      </software-slot>
    </software-inventory>
    <hardware xmlns="urn:ietf:params:xml:ns:yang:ietf-hardware">
      <component>
        <name>00_rrh</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-RADIO</class>
        <alias>rrh1608621390898</alias>
        <state>
          <admin-state>unlocked</admin-state>
          <oper-state>enabled</oper-state>
          <usage-state>unknown</usage-state>
        </state>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">00_rrh</o-ran-name>
        <physical-index>1</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev>%s</software-rev>
        <serial-num>%s</serial-num>
        <mfg-name>%s</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <uuid>D9A95B11-899A-4B5B-8119-A95BA9118199</uuid>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>01_fpga</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-FPGA</class>
        <alias>fpga</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">01_fpga</o-ran-name>
        <physical-index>2</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>02_board</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-BOARD</class>
        <alias>board</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">02_board</o-ran-name>
        <physical-index>3</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>03_power-ic</name>
        <class
          xmlns:ianahw="urn:ietf:params:xml:ns:yang:iana-hardware">ianahw:power-supply</class>
        <alias>power-ic</alias>
        <physical-index>4</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>04_radio-0</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-RADIO-0</class>
        <alias>radio-0</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">04_radio-0</o-ran-name>
        <physical-index>5</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>05_radio-1</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-RADIO-1</class>
        <alias>radio-1</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">05_radio-1</o-ran-name>
        <physical-index>6</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>06_radio-2</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-RADIO-2</class>
        <alias>radio-2</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">06_radio-2</o-ran-name>
        <physical-index>7</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>07_radio-3</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-RADIO-3</class>
        <alias>radio-3</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">07_radio-3</o-ran-name>
        <physical-index>8</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>08_transceiver-0</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-TRANSCEIVER-0</class>
        <alias>transceiver-0</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">08_transceiver-0</o-ran-name>
        <physical-index>9</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>09_transceiver-1</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-TRANSCEIVER-1</class>
        <alias>transceiver-1</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">09_transceiver-1</o-ran-name>
        <physical-index>10</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>10_transceiver-2</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-TRANSCEIVER-2</class>
        <alias>transceiver-2</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">10_transceiver-2</o-ran-name>
        <physical-index>11</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
      <component>
        <name>11_transceiver-3</name>
        <class
          xmlns:o-ran-hw="urn:o-ran:hardware:1.0">o-ran-hw:O-RAN-TRANSCEIVER-3</class>
        <alias>rrh</alias>
        <sensor-data>
          <value>0</value>
          <value-precision>0</value-precision>
          <oper-status>ok</oper-status>
          <units-display/>
          <value-update-rate>0</value-update-rate>
        </sensor-data>
        <label-content xmlns="urn:o-ran:hardware:1.0">
          <model-name>false</model-name>
          <serial-number>false</serial-number>
        </label-content>
        <o-ran-name xmlns="urn:o-ran:hardware:1.0">11_transceiver-3</o-ran-name>
        <physical-index>12</physical-index>
        <description/>
        <hardware-rev/>
        <firmware-rev/>
        <software-rev/>
        <serial-num/>
        <mfg-name>FACA</mfg-name>
        <model-name>RPQN-780x</model-name>
        <is-fru>false</is-fru>
        <product-code xmlns="urn:o-ran:hardware:1.0">5G-RHUB-V1</product-code>
      </component>
    </hardware>
    <operational-info xmlns="urn:o-ran:operations:1.0">
      <operational-state>
        <restart-cause>SUPERVISION-WATCHDOG</restart-cause>
      </operational-state>
    </operational-info>
    <interfaces xmlns="urn:ietf:params:xml:ns:yang:ietf-interfaces">
      <interface>
        <name>iC</name>
        <type
          xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:l2vlan</type>
        <enabled>true</enabled>
        <vlan-id xmlns="urn:o-ran:interfaces:1.0">1024</vlan-id>
        <mac-address xmlns="urn:o-ran:interfaces:1.0">6C:AD:AD:00:00:CC</mac-address>
        <port-reference xmlns="urn:o-ran:interfaces:1.0">
          <port-number>0</port-number>
        </port-reference>
        <admin-status>up</admin-status>
        <oper-status>up</oper-status>
        <last-change>1970-01-01T00:00:03Z</last-change>
        <if-index>1</if-index>
        <speed>10000</speed>
        <last-cleared xmlns="urn:o-ran:interfaces:1.0">1970-01-01T00:00:00Z</last-cleared>
      </interface>
      <interface>
        <name>iM</name>
        <type
          xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:ethernetCsmacd</type>
        <enabled>true</enabled>
        <statistics>
          <discontinuity-time>1970-01-01T00:00:03Z</discontinuity-time>
          <in-octets>6481408</in-octets>
          <in-unicast-pkts>94789</in-unicast-pkts>
          <in-broadcast-pkts>0</in-broadcast-pkts>
          <in-multicast-pkts>0</in-multicast-pkts>
          <in-discards>1</in-discards>
          <in-errors>0</in-errors>
          <in-unknown-protos>0</in-unknown-protos>
          <out-octets>2352440</out-octets>
          <out-unicast-pkts>36544</out-unicast-pkts>
          <out-broadcast-pkts>0</out-broadcast-pkts>
          <out-multicast-pkts>0</out-multicast-pkts>
          <out-discards>0</out-discards>
          <out-errors>0</out-errors>
        </statistics>
        <ipv4 xmlns="urn:ietf:params:xml:ns:yang:ietf-ip">
          <enabled>true</enabled>
          <address>
            <ip>10.240.10.2</ip>
            <netmask>255.255.255.0</netmask>
            <origin>dhcp</origin>
          </address>
        </ipv4>
        <l2-mtu xmlns="urn:o-ran:interfaces:1.0">1500</l2-mtu>
        <vlan-tagging xmlns="urn:o-ran:interfaces:1.0">true</vlan-tagging>
        <class-of-service xmlns="urn:o-ran:interfaces:1.0">
          <u-plane-marking>7</u-plane-marking>
          <c-plane-marking>7</c-plane-marking>
          <m-plane-marking>2</m-plane-marking>
          <s-plane-marking>7</s-plane-marking>
          <other-marking>1</other-marking>
        </class-of-service>
        <port-reference xmlns="urn:o-ran:interfaces:1.0">
          <port-number>2</port-number>
        </port-reference>
        <admin-status>up</admin-status>
        <oper-status>up</oper-status>
        <last-change>1970-01-01T00:00:03Z</last-change>
        <if-index>3</if-index>
        <phys-address>6c:ad:ad:00:00:00</phys-address>
        <speed>1000</speed>
        <last-cleared xmlns="urn:o-ran:interfaces:1.0">1970-01-01T00:00:00Z</last-cleared>
      </interface>
      <interface>
        <name>iU</name>
        <type
          xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:l2vlan</type>
        <enabled>true</enabled>
        <vlan-id xmlns="urn:o-ran:interfaces:1.0">1024</vlan-id>
        <mac-address xmlns="urn:o-ran:interfaces:1.0">6C:AD:AD:00:00:CC</mac-address>
        <port-reference xmlns="urn:o-ran:interfaces:1.0">
          <port-number>1</port-number>
        </port-reference>
        <admin-status>up</admin-status>
        <oper-status>up</oper-status>
        <last-change>1970-01-01T00:00:03Z</last-change>
        <if-index>2</if-index>
        <speed>10000</speed>
        <last-cleared xmlns="urn:o-ran:interfaces:1.0">1970-01-01T00:00:00Z</last-cleared>
      </interface>
<interface>
                <name>fheth0-no-VLAN</name>
                <type
                    xmlns:ianaift="urn:ietf:params:xml:ns:yang:iana-if-type">ianaift:l2vlan
                
                </type>
                <enabled>true</enabled>
                <oper-status>up</oper-status>
                <phys-address>d4:61:37:90:01:c1</phys-address>
                <lower-layer-if>fheth0</lower-layer-if>
                <ipv6
                    xmlns="urn:ietf:params:xml:ns:yang:ietf-ip">
                    <enabled>true</enabled>
                    <address>
                        <ip>fddc:f09a::106</ip>
                        <prefix-length>64</prefix-length>
                        <origin>dhcp</origin>
                    </address>
<address>
                        <ip>fe80:f09a::106</ip>
                        <prefix-length>64</prefix-length>
                        <origin>dhcp</origin>
                    </address>
                </ipv6>
                <mac-address
                    xmlns="urn:o-ran:interfaces:1.0">d4:61:37:90:01:c1
                
                </mac-address>
                <base-interface
                    xmlns="urn:o-ran:interfaces:1.0">fheth0
                
                </base-interface>
                <vlan-id
                    xmlns="urn:o-ran:interfaces:1.0">4
                
                </vlan-id>
            </interface>
    </interfaces>
    <module-capability xmlns="urn:o-ran:module-cap:1.0">
      <ru-capabilities>
        <ru-supported-category>CAT_A</ru-supported-category>
        <number-of-ru-ports>4</number-of-ru-ports>
        <number-of-spatial-streams>2</number-of-spatial-streams>
        <fronthaul-split-option>7</fronthaul-split-option>
        <format-of-iq-sample>
          <dynamic-compression-supported>true</dynamic-compression-supported>
          <realtime-variable-bit-width-supported>false</realtime-variable-bit-width-supported>
          <compression-method-supported>
            <compression-type>STATIC</compression-type>
            <bitwidth>8</bitwidth>
          </compression-method-supported>
          <compression-method-supported>
            <compression-type>STATIC</compression-type>
            <bitwidth>16</bitwidth>
          </compression-method-supported>
          <compression-method-supported>
            <compression-type>DYNAMIC</compression-type>
          </compression-method-supported>
        </format-of-iq-sample>
      </ru-capabilities>
<band-capabilities>
                <band-number>1</band-number>
                <max-supported-frequency-dl>1995000000</max-supported-frequency-dl>
                <min-supported-frequency-dl>1930000000</min-supported-frequency-dl>
                <max-supported-bandwidth-dl>65000000</max-supported-bandwidth-dl>
                <max-num-carriers-dl>2</max-num-carriers-dl>
                <max-carrier-bandwidth-dl>20000000</max-carrier-bandwidth-dl>
                <min-carrier-bandwidth-dl>5000000</min-carrier-bandwidth-dl>
                <max-supported-frequency-ul>1915000000</max-supported-frequency-ul>
                <min-supported-frequency-ul>1850000000</min-supported-frequency-ul>
                <max-supported-bandwidth-ul>65000000</max-supported-bandwidth-ul>
                <max-num-carriers-ul>2</max-num-carriers-ul>
                <max-carrier-bandwidth-ul>20000000</max-carrier-bandwidth-ul>
                <min-carrier-bandwidth-ul>5000000</min-carrier-bandwidth-ul>
                <max-num-bands>1</max-num-bands>
                <max-power-per-antenna>26.0</max-power-per-antenna>
                <min-power-per-antenna>0.0</min-power-per-antenna>
                <max-num-component-carriers>2</max-num-component-carriers>
                <max-num-sectors>1</max-num-sectors>
            </band-capabilities>
    </module-capability>
  </data>
</rpc-reply>`
)
