import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../../themes';

export default StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    borderRadius: 7,
  },
  buttonText: {
    fontSize: Scaled.fontSize.h10,
    fontFamily: Fonts.medium,
  },
});
