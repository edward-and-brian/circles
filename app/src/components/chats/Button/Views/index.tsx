import React, { PureComponent } from 'react';
import { Animated, TouchableOpacity, Text, View } from 'react-native';
import moment from 'moment';
import AvatarButton from '../../../shared/Button/AvatarButton';
import CircleWindow from '../../CircleWindow';
import { Scaled } from '../../../../themes';
import { Chat } from '../../ScrollList';
import { Date } from '../../../../utils';
import styles from './styles';

export interface Props {
  chat: Chat;
  circleWindowHeight: Animated.Value;
  onPressChat(): void;
  toggleCircleWindow(): void;
}

const date = moment().subtract(1, 'day');

class ButtonView extends PureComponent<Props> {
  render() {
    const {
      chat,
      onPressChat,
      toggleCircleWindow,
      circleWindowHeight,
    } = this.props;
    const mostRecentCircle = chat.circles[0];

    return (
      <View>
        <TouchableOpacity
          style={styles.chatContainer}
          onPress={onPressChat}
          onLongPress={toggleCircleWindow}
        >
          <View style={styles.avatarContainer}>
            <AvatarButton
              color={chat.avatar}
              diameter={Scaled.screen.height * 0.06}
            />
          </View>
          <View style={styles.textContainer}>
            <Text style={styles.name}>{chat.name}</Text>
            <View style={styles.circleAndDateContainer}>
              <Text style={styles.recentCircleName}>
                {mostRecentCircle.name}
              </Text>
              <Text style={styles.date}>
                {Date.toDisplayString(mostRecentCircle.lastMessage.createdAt)}
              </Text>
            </View>
          </View>
        </TouchableOpacity>
        <CircleWindow
          circles={chat.circles}
          circleWindowHeight={circleWindowHeight}
          onPressCircle={onPressChat}
        />
      </View>
    );
  }
}

export default ButtonView;
