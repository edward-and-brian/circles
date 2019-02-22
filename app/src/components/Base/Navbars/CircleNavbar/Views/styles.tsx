import { StyleSheet } from 'react-native';
import { Colors, Scaled } from '../../../../../themes';

export default StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'row',
    paddingTop: Scaled.navBarOffset,
    shadowOffset: { width: 0.5, height: 1 },
    shadowOpacity: 0.8,
    backgroundColor: Colors.white,
  },
  arrowContainer: {
    flex: 2,
    justifyContent: 'center',
    alignItems: 'center',
  },
  backArrow: {
    backgroundColor: Colors.black,
    width: 30,
    height: 30,
  },
  titleContainer: {
    flex: 9,
    justifyContent: 'center',
    paddingLeft: Scaled.screen.width * 0.05,
  },
  chatTitle: {
    fontSize: Scaled.fontSize.h11,
    color: Colors.gray,
  },
  circleTitle: {
    fontSize: Scaled.fontSize.h6,
  },
  avatarContainer: {
    flex: 2,
  },
});
