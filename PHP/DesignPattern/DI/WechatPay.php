<?php

/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
class WechatPay extends PayMethod {
    public function pay($amount) {
        return ('wechatpay paid:' . $amount);
    }
}