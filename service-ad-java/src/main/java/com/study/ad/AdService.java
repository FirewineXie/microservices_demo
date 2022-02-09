package com.study.ad;

import com.google.common.collect.ImmutableListMultimap;
import com.google.common.collect.Iterables;
import hipstershop.v1.DemoAd.Ad;
import io.grpc.Server;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.Random;

/**
 * @author Firewine Xie
 * @version 1.0.0
 * @ClassName AdService
 * @createTime: 2021年08月03日 14:34:18
 * @Description TODO
 */
public class AdService {

    private static final ImmutableListMultimap<String, Ad> adsMap = createAdsMap();

    @SuppressWarnings("FieldCanBeLocal")
    private static int MAX_ADS_TO_SERVE = 2;

    private static ImmutableListMultimap<String, Ad> createAdsMap() {
        Ad camera =
                Ad.newBuilder()
                        .setRedirectUrl("/product/2ZYFJ3GM2N")
                        .setText("Film camera for sale. 50% off.")
                        .build();
        Ad lens =
                Ad.newBuilder()
                        .setRedirectUrl("/product/66VCHSJNUP")
                        .setText("Vintage camera lens for sale. 20% off.")
                        .build();
        Ad recordPlayer =
                Ad.newBuilder()
                        .setRedirectUrl("/product/0PUK6V6EV0")
                        .setText("Vintage record player for sale. 30% off.")
                        .build();
        Ad bike =
                Ad.newBuilder()
                        .setRedirectUrl("/product/9SIQT8TOJO")
                        .setText("City Bike for sale. 10% off.")
                        .build();
        Ad baristaKit =
                Ad.newBuilder()
                        .setRedirectUrl("/product/1YMWWN1N4O")
                        .setText("Home Barista kitchen kit for sale. Buy one, get second kit for free")
                        .build();
        Ad airPlant =
                Ad.newBuilder()
                        .setRedirectUrl("/product/6E92ZMYYFZ")
                        .setText("Air plants for sale. Buy two, get third one for free")
                        .build();
        Ad terrarium =
                Ad.newBuilder()
                        .setRedirectUrl("/product/L9ECAV7KIM")
                        .setText("Terrarium for sale. Buy one, get second one for free")
                        .build();
        return ImmutableListMultimap.<String, Ad>builder()
                .putAll("photography", camera, lens)
                .putAll("vintage", camera, lens, recordPlayer)
                .put("cycling", bike)
                .put("cookware", baristaKit)
                .putAll("gardening", airPlant, terrarium)
                .build();
    }


    public Collection<Ad> getAdsByCategory(String category) {

        return adsMap.get(category);
    }

    private static final Random random = new Random();

    public List<Ad> getRandomAds() {
        List<Ad> ads = new ArrayList<>(MAX_ADS_TO_SERVE);
        Collection<Ad> allAds = adsMap.values();
        for (int i = 0; i < MAX_ADS_TO_SERVE; i++) {
            ads.add(Iterables.get(allAds, random.nextInt(allAds.size())));
        }
    }
}
