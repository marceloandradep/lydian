package draft

/*
// defines for materials, follow our polygon attributes as much as possible
#define MATV1_ATTR_2SIDED 0x0001
#define MATV1_ATTR_TRANSPARENT 0x0002
#define MATV1_ATTR_8BITCOLOR 0x0004
#define MATV1_ATTR_RGB16 0x0008
#define MATV1_ATTR_RGB24 0x0010
#define MATV1_ATTR_SHADE_MODE_CONSTANT 0x0020
#define MATV1_ATTR_SHADE_MODE_EMMISIVE 0x0020 // alias
#define MATV1_ATTR_SHADE_MODE_FLAT 0x0040
#define MATV1_ATTR_SHADE_MODE_GOURAUD 0x0080
#define MATV1_ATTR_SHADE_MODE_FASTPHONG 0x0100
#define MATV1_ATTR_SHADE_MODE_TEXTURE 0x0200

// states of materials
#define MATV1_STATE_ACTIVE 0x0001

// defines for material system
#define MAX_MATERIALS 256
*/

/*
// a first version of a “material”
typedef struct MATV1_TYP
{
	int state; 				// state of material
	int id; 				// id of this material, index into material array
	char name[64]; 			// name of material
	int attr; 				// attributes, the modes for shading, constant, flat,
							// gouraud, fast phong, environment, textured etc.
							// and other special flags...
	RGBAV1 color; 			// color of material
	float ka, kd, ks, power;	// ambient, diffuse, specular,
								// coefficients, note they are
								// separate and scalars since many
								// modelers use this format
								// along with specular power
	RGBAV1 ra, rd, rs; 		// the reflectivities/colors pre-
							// multiplied, to more match our
							// definitions, each is basically
							// computed by multiplying the
							// color by the k’s, eg:
							// rd = color*kd etc.
	char texture_file[80];	// file location of texture
	BITMAP texture;			// actual texture map (if any)
} MATV1, *MATV1_PTR;
*/
