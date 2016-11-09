<?php

/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
class AliPay extends PayMethod {

    /**
     * AliPay constructor.
     */
    public function __construct() {
        //var_dump('call alipay construct');
    }

    public function pay($amount) {
       return ('alipay paid:' . $amount);
    }
}