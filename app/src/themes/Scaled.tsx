import { Dimensions } from 'react-native';

const { width: SCREEN_WIDTH, height: SCREEN_HEIGHT } = Dimensions.get('window');
const NORMALIZED_FONTS = SCREEN_WIDTH * 0.024;
const isXGen = SCREEN_HEIGHT > 800;

const Scaled = {
  screen: {
    width: SCREEN_WIDTH,
    height: SCREEN_HEIGHT,
  },
  fontSize: {
    h1: NORMALIZED_FONTS * 7.6,
    h2: NORMALIZED_FONTS * 5.1,
    h3: NORMALIZED_FONTS * 3.8,
    h4: NORMALIZED_FONTS * 3.33,
    h5: NORMALIZED_FONTS * 2.67,
    h6: NORMALIZED_FONTS * 2.45,
    h7: NORMALIZED_FONTS * 2.22,
    h8: NORMALIZED_FONTS * 2,
    h9: NORMALIZED_FONTS * 1.77,
    h10: NORMALIZED_FONTS * 1.55,
    h11: NORMALIZED_FONTS * 1.34,
    h12: NORMALIZED_FONTS * 1.1,
    h13: NORMALIZED_FONTS,
  },
  isXGen,
  navBarOffset: isXGen ? SCREEN_WIDTH * 0.12 : SCREEN_WIDTH * 0.05,
};

export default Scaled;
