import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  container: {
    flex: 1,
    width: '100%',
    alignItems: 'center',
  },
  label: {
    fontFamily: Fonts.medium,
    fontSize: Scaled.fontSize.h11,
  },
  input: {
    width: '95%',
    justifyContent: 'center',
    height: Scaled.screen.height * 0.045,
    paddingLeft: Scaled.screen.height * 0.013,
    borderRadius: Scaled.screen.height * 0.008,
    fontFamily: Fonts.book,
    fontSize: Scaled.fontSize.h10,
    backgroundColor: Colors.lightGray,
  },
});
