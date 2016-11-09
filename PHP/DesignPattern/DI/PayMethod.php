<?php

/**
 * 支付抽象类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
abstract class PayMethod {
    abstract function pay($amount);
}