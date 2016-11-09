<?php

/**
 * 具体建造者类A
 * 实现抽象建造者类
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/7/28
 */
class SubMealBuilderA extends MealBuilder {
    private $meal;

    public function __construct() {
        $this->meal = new Meal();//具体建造的产品对象
    }

    function BuildDrink() {
        $this->meal->setDrink('build drink A');
    }

    function BuildFood() {
        $this->meal->setFood('build food A');
    }

    function getMeal() {
        return $this->meal->getDrink() . '+++' . $this->meal->getFood();
    }

}