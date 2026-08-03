//-----------------------------------------------------------------------------
/*

Array88 Driver

*/
//-----------------------------------------------------------------------------

#include "array88.h"

//-----------------------------------------------------------------------------

#define NUM_ROWS 8
#define NUM_COLS 8

#define RED (1 << 0)
#define GREEN (1 << 1)
#define BLUE (1 << 2)

static uint8_t red[NUM_ROWS];
static uint8_t green[NUM_ROWS];
static uint8_t blue[NUM_ROWS];

//-----------------------------------------------------------------------------

__sfr __at 0x05 y88Port;	// Standard 8x8 Row (Y) select
__sfr __at 0x06 xr88Port;	// RGB 8x8 (Red) column (X) select
__sfr __at 0xf8 xg88Port;	// RGB 8x8 (Green) column (X) select
__sfr __at 0xf9 xb88Port;	// RGB 8x8 (Blue) column (X) select

//-----------------------------------------------------------------------------

static void delay(void) {
	for (volatile int8_t i = 0; i < 30; i++) ;
}

void array88_off(void) {
	for (int8_t i = 0; i < NUM_ROWS; i++) {
		red[i] = 0;
		green[i] = 0;
		blue[i] = 0;
	}
}

void array88_scan(void) {
	for (int8_t i = 0; i < 8; i++) {
		xr88Port = red[i];
		xg88Port = green[i];
		xb88Port = blue[i];
		y88Port = 1 << i;
		delay();
		y88Port = 0;
	}
}

void array88_plot(uint8_t x, uint8_t y, uint8_t color) {
	x &= (NUM_COLS - 1);
	y &= (NUM_ROWS - 1);

	uint8_t xmask = 1 << x;

	if ((color & RED) != 0) {
		red[y] |= xmask;
	} else {
		red[y] &= ~xmask;
	}

	if ((color & GREEN) != 0) {
		green[y] |= xmask;
	} else {
		green[y] &= ~xmask;
	}

	if ((color & BLUE) != 0) {
		blue[y] |= xmask;
	} else {
		blue[y] &= ~xmask;
	}
}

//-----------------------------------------------------------------------------
