<?php

/**容器类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/4
 */
class DI {
    private $obj_list_arr = array();// 容器存放数组
    protected static $instince = null;

    public static function getInstince() {// 单例模式
        if (self::$instince == null) {
            self::$instince = new DI();
        }
        return self::$instince;
    }

    function __get($name) {// 魔术方法，用于匿名赋值和调用
        return $this->obj_list_arr[$name];
    }

    function __set($name, $value) {
        $this->obj_list_arr[$name] = $value;
    }
}