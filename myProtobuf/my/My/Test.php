<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: my.proto

namespace My;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Protobuf type <code>my.Test</code>
 */
class Test extends \Google\Protobuf\Internal\Message
{
    /**
     * <code>string label = 1;</code>
     */
    private $label = '';
    /**
     * <code>int32 type = 2;</code>
     */
    private $type = 0;
    /**
     * <code>repeated int64 reps = 3;</code>
     */
    private $reps;
    /**
     * <code>string requiredField = 5;</code>
     */
    private $requiredField = '';

    public function __construct() {
        \GPBMetadata\My::initOnce();
        parent::__construct();
    }

    /**
     * <code>string label = 1;</code>
     */
    public function getLabel()
    {
        return $this->label;
    }

    /**
     * <code>string label = 1;</code>
     */
    public function setLabel($var)
    {
        GPBUtil::checkString($var, True);
        $this->label = $var;
    }

    /**
     * <code>int32 type = 2;</code>
     */
    public function getType()
    {
        return $this->type;
    }

    /**
     * <code>int32 type = 2;</code>
     */
    public function setType($var)
    {
        GPBUtil::checkInt32($var);
        $this->type = $var;
    }

    /**
     * <code>repeated int64 reps = 3;</code>
     */
    public function getReps()
    {
        return $this->reps;
    }

    /**
     * <code>repeated int64 reps = 3;</code>
     */
    public function setReps(&$var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::INT64);
        $this->reps = $arr;
    }

    /**
     * <code>string requiredField = 5;</code>
     */
    public function getRequiredField()
    {
        return $this->requiredField;
    }

    /**
     * <code>string requiredField = 5;</code>
     */
    public function setRequiredField($var)
    {
        GPBUtil::checkString($var, True);
        $this->requiredField = $var;
    }

}

