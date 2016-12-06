<?php
/**
 * Created by DannyWang
 * wangjue@mia.com
 * 2016/8/15
 */

class Temp{
    private $type;
    private $val;

    /**
     * Temp constructor.
     * @param $type
     */
    public function __construct($type) {
        $this->type = $type;
    }

    /**
     * @return mixed
     */
    public function getVal() {
        return $this->val;
    }

    /**
     * @param mixed $val
     */
    public function setVal($val) {
        $this->val = $val;
        return $this;
    }

    public function getString(){
        //return "type:{$this->type} val:{$this->val}";
    }


}

echo (new Temp('item'))->setVal('vvvvv')->getVal();
echo (new Temp('item'))->setVal('valval')->getString();
