<?php

/**
 * 路 抽象类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/1
 */
abstract class Road {
    // 路上行驶的汽车
    protected $car;

    // 开始走
    abstract public function Go();

    // 设置路上的汽车（组合）
    public function setCar(Car $car) {
        $this->car = $car;
    }
}