import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../../themes';

export default StyleSheet.create({
  container: {
    flexDirection: 'row',
    paddingTop: Scaled.navBarOffset,
    shadowOffset: { width: 0.5, height: 1 },
    shadowOpacity: 0.8,
    backgroundColor: Colors.white,
  },
  arrowContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  backArrow: {
    width: Scaled.screen.width * 0.065,
    height: Scaled.screen.width * 0.065,
  },
  titleContainer: {
    flex: 5,
    justifyContent: 'center',
  },
  chatTitle: {
    fontSize: Scaled.fontSize.h11,
    fontFamily: Fonts.medium,
    color: Colors.gray,
  },
  circleTitle: {
    fontSize: Scaled.fontSize.h6,
    fontFamily: Fonts.medium,
  },
  avatarContainer: {
    flex: 1,
    justifyContent: 'center',
  },
});
