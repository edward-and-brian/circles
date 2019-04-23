import React, { PureComponent } from 'react';
import { TouchableOpacity, Text, View } from 'react-native';
import AvatarButton from '../../../shared/Button/AvatarButton';
import { Scaled } from '../../../../themes';
import { Circle } from '../../ScrollList';
import { Date } from '../../../../utils';
import styles from './styles';

export interface Props {
  circle: Circle;
  onPressCircle(): void;
}

class CircleButtonView extends PureComponent<Props> {
  render() {
    const { circle, onPressCircle } = this.props;
    let { content } = circle.lastMessage;
    return (
      <TouchableOpacity onPress={onPressCircle} style={styles.circleButton}>
        <View style={styles.avatarContainer}>
          <AvatarButton
            color={circle.avatar}
            diameter={Scaled.screen.height * 0.04}
          />
        </View>
        <View style={styles.circleTextContainer}>
          <Text style={styles.circleName}>{circle.name}</Text>
          <View style={styles.recentMessageAndDateContainer}>
            <Text style={styles.recentMessageText} numberOfLines={1}>
              {content}
            </Text>
            <Text style={styles.recentMessageDate}>
              {Date.toDisplayString(circle.lastMessage.createdAt)}
            </Text>
          </View>
        </View>
      </TouchableOpacity>
    );
  }
}

export default CircleButtonView;
