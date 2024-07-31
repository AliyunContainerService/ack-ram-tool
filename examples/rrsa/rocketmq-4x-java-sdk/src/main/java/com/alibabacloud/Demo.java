package com.alibabacloud;

// com.aliyun:credentials-java >= 0.3.5
import com.aliyun.credentials.Client;

import com.aliyun.credentials.models.CredentialModel;
import org.apache.rocketmq.acl.common.AclClientRPCHook;
import org.apache.rocketmq.acl.common.SessionCredentials;
import org.apache.rocketmq.client.producer.DefaultMQProducer;
import org.apache.rocketmq.client.producer.SendResult;
import org.apache.rocketmq.client.consumer.DefaultMQPushConsumer;
import org.apache.rocketmq.client.consumer.listener.ConsumeConcurrentlyContext;
import org.apache.rocketmq.client.consumer.listener.ConsumeConcurrentlyStatus;
import org.apache.rocketmq.client.consumer.listener.MessageListenerConcurrently;
import org.apache.rocketmq.common.message.Message;
import org.apache.rocketmq.common.message.MessageExt;
import org.apache.rocketmq.remoting.RPCHook;
import org.apache.rocketmq.remoting.common.RemotingHelper;
import org.apache.rocketmq.remoting.protocol.RemotingCommand;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;
import java.util.Date;
import java.util.List;


class CustomAclClientRPCHook implements RPCHook {
    private final com.aliyun.credentials.Client cred;

    public CustomAclClientRPCHook(com.aliyun.credentials.Client cred) {
        this.cred = cred;
    }

    public void doBeforeRequest(String remoteAddr, RemotingCommand request) {
        SessionCredentials sc = this.getSessionCredentials();
        AclClientRPCHook hook = new AclClientRPCHook(sc);
        hook.doBeforeRequest(remoteAddr, request);
    }

    public void doAfterResponse(String remoteAddr, RemotingCommand request, RemotingCommand response) {
    }

    synchronized public SessionCredentials getSessionCredentials() {
        CredentialModel cm = cred.getCredential();
        String ak = cm.getAccessKeyId();
        String sk = cm.getAccessKeySecret();
        String token = cm.getSecurityToken();
        if (token != null) {
            return new SessionCredentials(ak, sk, token);
        } else {
            return new SessionCredentials(ak, sk);
        }
    }
}

class TestMQ4XSDK {

    private static RPCHook getAclRPCHook(com.aliyun.credentials.Client cred) {
        return new CustomAclClientRPCHook(cred);
    }

    public void ProducerExample(com.aliyun.credentials.Client cred) throws Exception {
        DefaultMQProducer producer = new DefaultMQProducer(getAclRPCHook(cred));

        // producer.setAccessChannel(AccessChannel.CLOUD);
        // MQ_INST_XX.XX.mq.aliyuncs.com:80
        producer.setNamesrvAddr(System.getenv("MQ_ENDPOINT"));
        producer.start();

        try {
            Message msg = new Message("yourNormalTopic",
                    "YOUR MESSAGE TAG",
                    "Hello world".getBytes(RemotingHelper.DEFAULT_CHARSET));
            SendResult sendResult = producer.send(msg);
            System.out.printf("[Producer][%s] SendResult msgId: %s\n", now(), sendResult.getMsgId());
        } catch (Exception e) {
            System.out.println(now() + " Send mq message failed.");
            e.printStackTrace();
        }

        producer.shutdown();
    }

    public void ConsumerExample(com.aliyun.credentials.Client cred) throws Exception {
        DefaultMQPushConsumer consumer = new DefaultMQPushConsumer(getAclRPCHook(cred));

        // consumer.setAccessChannel(AccessChannel.CLOUD);
        // MQ_INST_XX.XX.mq.aliyuncs.com:80
        consumer.setNamesrvAddr(System.getenv("MQ_ENDPOINT"));

        consumer.setConsumerGroup("yourConsumerGroup");
        consumer.subscribe("yourNormalTopic", "*");
//        consumer.setConsumeTimestamp("20180422221800");
        consumer.registerMessageListener(new MessageListenerConcurrently() {
            @Override
            public ConsumeConcurrentlyStatus consumeMessage(List<MessageExt> msgs, ConsumeConcurrentlyContext context) {
                System.out.printf("[Consumer][%s] Receive New Messages: [%s]\n", now(), msgIds(msgs));
                return ConsumeConcurrentlyStatus.CONSUME_SUCCESS;
            }
        });
        consumer.start();
        System.out.printf("Consumer Started.%n");
    }

    private static String msgIds(List<MessageExt> msg) {
        StringBuilder msgIds = new StringBuilder();
        if (msg == null || msg.size() == 0)
            return msgIds.toString();
        for (MessageExt m : msg) {
            if (m != null) {
                msgIds.append("msgId: ").append(m.getMsgId()).append(", ");
            }
        }
        return msgIds.toString();
    }

    private static String now() {
        return LocalDateTime.now().format(DateTimeFormatter.ofPattern("yyyy-MM-dd HH:mm:ss"));
    }
}


public class Demo {
    public static void main(String[] args) throws Exception {
        // 两种方式都可以
        com.aliyun.credentials.Client cred = new com.aliyun.credentials.Client();
        // or
        // com.aliyun.credentials.Client cred = newOidcCred();

        // test RocketMQ 4.x sdk (https://github.com/apache/rocketmq/tree/rocketmq-all-4.9.8) use rrsa oidc token
        System.out.println("test RocketMQ 4.x sdk use rrsa oidc token");
        TestMQ4XSDK example = new TestMQ4XSDK();

        System.out.println("====== Producer =======");
        example.ProducerExample(cred);

        System.out.println("====== Consumer =======");
        example.ConsumerExample(cred);
    }

    static com.aliyun.credentials.Client newOidcCred() throws Exception {
        // new credential which use rrsa oidc token
        com.aliyun.credentials.models.Config credConf = new com.aliyun.credentials.models.Config();
        credConf.type = "oidc_role_arn";
        credConf.roleArn = System.getenv("ALIBABA_CLOUD_ROLE_ARN");
        credConf.oidcProviderArn = System.getenv("ALIBABA_CLOUD_OIDC_PROVIDER_ARN");
        credConf.oidcTokenFilePath = System.getenv("ALIBABA_CLOUD_OIDC_TOKEN_FILE");
        credConf.roleSessionName = "test-rrsa-oidc-token";
        // https://next.api.aliyun.com/product/Sts
        // credConf.setSTSEndpoint("sts-vpc.cn-hangzhou.aliyuncs.com")
        return new com.aliyun.credentials.Client(credConf);
    }
}
