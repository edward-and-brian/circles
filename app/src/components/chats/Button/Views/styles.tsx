import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  container: {
    height: Scaled.screen.height * 0.1,
    backgroundColor: Colors.white,
    margin: Scaled.screen.height * 0.006,
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: Scaled.screen.height * 0.01,
  },
  name: {
    fontSize: Scaled.fontSize.h4,
    fontFamily: Fonts.heavy,
  },
});
