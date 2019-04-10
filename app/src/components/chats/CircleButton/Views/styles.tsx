import { StyleSheet } from 'react-native';
import { Colors, Scaled, Fonts } from '../../../../themes';

export default StyleSheet.create({
  avatarContainer: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
  },
  circleButton: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'center',
    margin: Scaled.screen.height * 0.006,
    borderRadius: Scaled.screen.height * 0.01,
  },
  circleTextContainer: {
    flex: 5,
    justifyContent: 'center',
    marginVertical: Scaled.screen.height * 0.008,
  },
  circleName: {
    fontSize: Scaled.fontSize.h9,
    fontFamily: Fonts.medium,
  },
  recentMessageAndDateContainer: {
    flexDirection: 'row',
    justifyContent: 'space-between',
    alignItems: 'center',
  },
  recentMessageText: {
    fontSize: Scaled.fontSize.h11,
    fontFamily: Fonts.medium,
    color: Colors.gray,
  },
  recentMessageDate: {
    fontSize: Scaled.fontSize.h11,
    fontFamily: Fonts.medium,
  },
});
