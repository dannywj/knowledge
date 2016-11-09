<?php

/**
 * 高速公路类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/1
 */
class SpeedRoad extends Road {

    public function Go() {
        // 汽车启动
        $this->car->Run();
        var_dump('on SpeedRoad');
    }
}