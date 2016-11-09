<?php

/**
 * 街道类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/1
 */
class StreetRoad extends Road {
    public function Go() {
        $this->car->Run();
        var_dump('on StreetRoad');
    }
}