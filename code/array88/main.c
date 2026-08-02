//-----------------------------------------------------------------------------
/*

Array88 Driver

*/
//-----------------------------------------------------------------------------

__sfr __at 0x05 y88Port;	// Standard 8x8 Row (Y) select
__sfr __at 0x06 xr88Port;	// RGB 8x8 (Red) column (X) select
__sfr __at 0xf8 xg88Port;	// RGB 8x8 (Green) column (X) select
__sfr __at 0xf9 xb88Port;	// RGB 8x8 (Blue) column (X) select

__sfr __at 0xaa testPort;

//-----------------------------------------------------------------------------

#define NUM_ROWS 8
#define NUM_COLS 8

#define RED (1 << 0)
#define GREEN (1 << 1)
#define BLUE (1 << 2)

typedef unsigned char uint8_t;
typedef unsigned int uint16_t;

unsigned char red[NUM_ROWS];
unsigned char green[NUM_ROWS];
unsigned char blue[NUM_ROWS];

unsigned int seed = 12345;	// Seed

unsigned int rng(void) {
	// Multiplier (a) = 25173, Increment (c) = 13849
	seed = (seed * 25173) + 13849;
	return seed;
}

void delay() {
	for (volatile int i = 0; i < 30; i++) ;
}

void scan(void) {
	for (int i = 0; i < 8; i++) {
		xr88Port = red[i];
		xg88Port = green[i];
		xb88Port = blue[i];
		y88Port = 1 << i;
		delay();
		y88Port = 0;
		xr88Port = 0;
		xg88Port = 0;
		xb88Port = 0;
	}
}

void array88_plot(uint8_t x, uint8_t y, uint8_t color) {
	x &= (NUM_COLS - 1);
	y &= (NUM_ROWS - 1);

	uint8_t xmask = 1 << x;

	red[y] &= ~xmask;
	green[y] &= ~xmask;
	blue[y] &= ~xmask;

	if ((color & RED) != 0) {
		red[y] |= xmask;
	}
	if ((color & GREEN) != 0) {
		green[y] |= xmask;
	}
	if ((color & BLUE) != 0) {
		blue[y] |= xmask;
	}
}

void randomize(void) {
	for (int i = 0; i < 8; i++) {
		red[i] = rng();
		green[i] = rng();
		blue[i] = rng();
	}
}

void array88_off(void) {
	for (int i = 0; i < NUM_ROWS; i++) {
		red[i] = 0;
		green[i] = 0;
		blue[i] = 0;
	}
}

//-----------------------------------------------------------------------------

int main(void) {

	array88_off();

	for (int i = 0; i < 8; i++) {
		array88_plot(0, i, i);
		array88_plot(i, 0, i);
		array88_plot(i, i, i);
	}

	while (1) {
		scan();
	}

	//__asm__("halt");
	//return 0;

}

//-----------------------------------------------------------------------------
