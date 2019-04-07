import { StyleSheet } from 'react-native';
import { Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  container: {
    flexDirection: 'row',
    paddingTop: Scaled.navbarOffset,
    height: Scaled.navbarHeight,
  },
  titleContainer: {
    paddingLeft: Scaled.screen.width * 0.025,
    justifyContent: 'center',
    flex: 5,
  },
  title: {
    fontSize: Scaled.fontSize.h3,
    fontFamily: Fonts.heavy,
  },
  avatarContainer: {
    flex: 1,
    justifyContent: 'center',
  },
});
