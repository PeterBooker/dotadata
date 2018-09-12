<?php
/**
 * Primary class file for the DotaData plugin.
 *
 * @package dotadata
 */

/**
 * Class DotaData
 */
class DotaData {
	/**
	 * Notices to show at the head of the admin screen.
	 *
	 * @access public
	 *
	 * @var array
	 */
	public $admin_notices = array();

	/**
	 * DotaData constructor.
	 *
	 * @uses DotaData::init()
	 *
	 * @return void
	 */
	public function __construct() {
		$this->init();
	}

	/**
	 * Plugin initiation.
	 *
	 * A helper function, called by `DotaData::__construct()` to initiate actions, hooks and other features needed.
	 *
	 * @uses add_action()
	 * @uses add_filter()
	 *
	 * @return void
	 */
	public function init() {

	}
}
