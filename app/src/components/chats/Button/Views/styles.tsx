import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  chatContainer: {
    flexDirection: 'row',
    height: Scaled.screen.height * 0.1,
    marginHorizontal: Scaled.screen.height * 0.006,
    borderRadius: Scaled.screen.height * 0.01,
  },
  avatarContainer: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  textContainer: {
    flex: 5,
    justifyContent: 'center',
    marginVertical: Scaled.screen.height * 0.03,
  },
  name: {
    fontSize: Scaled.fontSize.h7,
    fontFamily: Fonts.medium,
  },
  circleAndDateContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  recentCircleName: {
    fontSize: Scaled.fontSize.h9,
    fontFamily: Fonts.medium,
    color: Colors.gray,
  },
  date: {
    fontSize: Scaled.fontSize.h10,
    fontFamily: Fonts.medium,
  },
  circleContainer: {
    marginLeft: '16.66%',
  },
  circleButton: {
    alignItems: 'center',
    justifyContent: 'center',
    height: Scaled.screen.height * 0.06,
    margin: Scaled.screen.height * 0.006,
    borderRadius: Scaled.screen.height * 0.01,
    backgroundColor: 'lightpink',
  },
});
