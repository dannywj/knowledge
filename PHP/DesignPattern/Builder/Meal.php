<?php

/**
 * 产品类
 * 具体建造的对象类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
class Meal {
    private $food;
    private $drink;

    /**
     * @return mixed
     */
    public function getFood() {
        return $this->food;
    }

    /**
     * @param mixed $food
     */
    public function setFood($food) {
        $this->food = $food;
    }

    /**
     * @return mixed
     */
    public function getDrink() {
        return $this->drink;
    }

    /**
     * @param mixed $drink
     */
    public function setDrink($drink) {
        $this->drink = $drink;
    }
}