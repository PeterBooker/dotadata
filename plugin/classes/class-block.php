<?php
/**
 * Class file for the Block.
 *
 * @package dotadata
 */

namespace DotaData;

/**
 * Class ShortCode
 */
class Block {

	/**
	 * Block constructor.
	 *
	 * @uses Block::init()
	 *
	 * @return void
	 */
	public function __construct() {
		$this->init();
	}

	/**
	 * Block initiation.
	 *
	 * @uses add_shortcode()
	 *
	 * @return void
	 */
	public function init() {
		register_block_type( 'dotadata/datasheet', array(
			'editor_script'   => 'dotadata-script',
			'editor_style'    => 'dotadata-style',
			'render_callback' => array( $this, 'render' ),
		) );

		add_action( 'enqueue_block_editor_assets', array( $this, 'editor_assets' ) );
	}

	/**
	 * Block Editor Assets.
	 *
	 * Enqueues Gutenberg Block editor assets.
	 *
	 * @return void
	 */
	public function editor_assets() {

		$dependencies = array(
			'wp-blocks',
			'wp-i18n',
			'wp-element',
			'wp-components',
		);

		wp_enqueue_script(
			'dotadata-script',
			DOTADATA_URL . 'assets/js/block.min.js',
			$dependencies,
			filemtime( DOTADATA_PATH . '/assets/js/block.min.js' )
		);

		wp_enqueue_style(
			'dotadata-style',
			DOTADATA_URL . 'assets/css/block.css',
			array(),
			filemtime( DOTADATA_PATH . '/assets/css/block.css' )
		);

	}

	/**
	 * Block Render.
	 *
	 * Renders the Block on the frontend.
	 *
	 * @param array $atts Array of attributes.
	 *
	 * @return string
	 */
	public function render( $atts ) {

		$atts = array_change_key_case( (array) $atts, CASE_LOWER );

		$atts = shortcode_atts( [
			'lang'     => 'en',
			'theme'    => 'default',
			'showlore' => true,
			'type'     => 'hero',
			'item'     => 'riki',
			'data'     => '',
		], $atts, $tag );

		$classes = array(
			'wp-block-dd-datasheet',
			'dotadata',
			'theme-' . $theme,
		);

		if ( ! empty( $data ) ) {

			wp_enqueue_style( 'dotadata-themes', DOTADATA_URL . 'assets/css/themes.css', array(), DOTADATA_VERSION, 'all' );

			return sprintf(
				'<div class="%1$s">%2$s</div>',
				join( ' ', $classes ),
				$highlighted,
				$styleblock
			);

		} else {

			return sprintf(
				'<div class="%1$s"><pre><code>%2$s</pre></code></div>',
				join( ' ', $classes ),
				esc_html( $content )
			);

		}

	}

}
