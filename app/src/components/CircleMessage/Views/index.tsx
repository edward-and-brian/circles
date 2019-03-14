import React, { PureComponent } from 'react';
import { View, Text } from 'react-native';
import moment from 'moment';
import CircleAvatarButton from '../../shared/Avatar/CircleAvatarButton';
import { Scaled } from '../../../themes';
import styles from './styles';

export interface Props {
  messageGroup: {
    user: {
      name: string;
      avatar: string;
    };
    messages: string[];
    date: Date;
  };
}

class CircleMessageView extends PureComponent<Props> {
  render() {
    const { messageGroup } = this.props;

    return (
      <View style={styles.container}>
        <View style={styles.avatarContainer}>
          <CircleAvatarButton
            color={messageGroup.user.avatar}
            diameter={Scaled.screen.width * 0.105}
          />
        </View>
        <View style={styles.contentContainer}>
          <View style={styles.nameAndDateContainer}>
            <Text style={styles.name}>{messageGroup.user.name}</Text>
            <Text style={styles.date}>
              {moment(messageGroup.date).format('h:mm a')}
            </Text>
          </View>
          {messageGroup.messages.reverse().map((message, index) => (
            <Text
              style={styles.message}
              key={`${messageGroup.user.name}${messageGroup.date}${index}`}
            >
              {message}
            </Text>
          ))}
        </View>
      </View>
    );
  }
}

export default CircleMessageView;
