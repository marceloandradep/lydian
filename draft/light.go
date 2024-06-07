package draft

/*
// defines for light types
#define LIGHTV1_ATTR_AMBIENT 0x0001 // basic ambient light
#define LIGHTV1_ATTR_INFINITE 0x0002 // infinite light source
#define LIGHTV1_ATTR_POINT 0x0004 // point light source
#define LIGHTV1_ATTR_SPOTLIGHT1 0x0008 // spotlight type 1 (simple)
#define LIGHTV1_ATTR_SPOTLIGHT2 0x0010 // spotlight type 2 (complex)
#define LIGHTV1_STATE_ON 1 // light on
#define LIGHTV1_STATE_OFF 0 // light off
#define MAX_LIGHTS 8 // good luck with 1!
*/

/*
// version 1.0 light structure
typedef struct LIGHTV1_TYP
{
int state; // state of light
int id; // id of light
int attr; // type of light, and extra qualifiers
RGBAV1 c_ambient; // ambient light intensity
RGBAV1 c_diffuse; // diffuse light intensity
RGBAV1 c_specular; // specular light intensity
POINT4D pos; // position of light
VECTOR4D dir; // direction of light
float kc, kl, kq; // attenuation factors
float spot_inner; // inner angle for spot light
float spot_outer; // outer angle for spot light
float pf; // power factor/falloff for spot lights
} LIGHTV1, *LIGHTV1_PTR;
*/
