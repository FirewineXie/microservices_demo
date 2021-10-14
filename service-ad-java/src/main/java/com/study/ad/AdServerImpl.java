package com.study.ad;

import hipstershop.v1.AdServiceGrpc;
import hipstershop.v1.DemoAd;
import io.grpc.stub.StreamObserver;
import org.apache.logging.log4j.Level;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;

/**
 * @author Firewine Xie
 * @version 1.0.0
 * @ClassName AdServerImpl
 * @createTime: 2021年08月03日 14:29:58
 * @Description TODO
 */
public class AdServerImpl extends AdServiceGrpc.AdServiceImplBase {


    private static final Logger logger = LogManager.getLogger(AdService.class);
    private AdService service;

    private AdService initService() {
        if (service != null) {
            return service;
        }
        service = new AdService();
        return service;
    }

    @Override
    public void getAds(DemoAd.AdRequest request, StreamObserver<DemoAd.AdResponse> responseObserver) {
        initService();

        List<DemoAd.Ad> allAds = new ArrayList<>();
        logger.info("received ad request (context_words = " + request.getContextKeysList() + ")");
        if (request.getContextKeysCount() > 0) {
            for (int i = 0; i < request.getContextKeysCount(); i++) {
                Collection<DemoAd.Ad> ads = service.getAdsByCategory(request.getContextKeys(i));
                allAds.addAll(ads);
            }
        }else{
            allAds = service.getRandomAds();
        }
        if (allAds.isEmpty()){
            allAds = service.getRandomAds();
        }
        DemoAd.AdResponse response = DemoAd.AdResponse.newBuilder().addAllAds(allAds).build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }
}
