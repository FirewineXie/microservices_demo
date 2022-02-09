package com.study.ad;

import io.grpc.Server;
import io.grpc.ServerBuilder;

import java.io.IOException;
import java.util.logging.Logger;

/**
 * @author Firewine Xie
 * @version 1.0.0
 * @ClassName AdServerClient
 * @createTime: 2021年08月03日 14:33:24
 * @Description TODO
 */
public class AdServerClient {

    private static final Logger logger = Logger.getLogger(AdServerClient.class.getName());

    private Server server;



    private void start() throws IOException {
        int port = 50051;
        server = ServerBuilder.forPort(port)
                .addService(new AdServerImpl())
                .build()
                .start();

        logger.info("Server started , listening on " + port);
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            logger.info("*** shutting down gRPC server since JVM is shutting down");
            AdServerClient.this.stop();
            logger.info("*** server shut down");
        }));
    }


    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    /**
     * Await termination on the main thread since the grpc library uses daemon threads.
     */
    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    /**
     * Main launches the server from the command line.
     */
    public static void main(String[] args) throws IOException, InterruptedException {
        final AdServerClient server = new AdServerClient();
        server.start();
        server.blockUntilShutdown();
    }
}
