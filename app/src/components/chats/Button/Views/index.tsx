import React, { PureComponent } from 'react';
import {
  Animated,
  TouchableOpacity,
  Text,
  View,
  ScrollView,
} from 'react-native';
import moment from 'moment';
import AvatarButton from '../../../shared/Button/AvatarButton';
import { Colors, Scaled } from '../../../../themes';
import { Chat } from '../../ScrollList';
import styles from './styles';

export interface Props {
  chat: Chat;
  circleWindowHeight: Animated.Value;
  onPressChat(): void;
  toggleCircleWindow(): void;
}

const date = moment().subtract(1, 'day');

class ButtonView extends PureComponent<Props> {
  renderCircleWindow() {
    return (
      <Animated.View
        style={[
          styles.circleContainer,
          { height: this.props.circleWindowHeight },
        ]}
      >
        <ScrollView>
          {this.props.chat.circles.map(circle => (
            <View style={styles.circleButton}>
              <Text>{circle}</Text>
            </View>
          ))}
        </ScrollView>
      </Animated.View>
    );
  }

  render() {
    return (
      <View>
        <TouchableOpacity
          style={styles.chatContainer}
          onPress={this.props.onPressChat}
          onLongPress={this.props.toggleCircleWindow}
        >
          <View style={styles.avatarContainer}>
            <AvatarButton
              color={Colors.red}
              diameter={Scaled.screen.height * 0.06}
            />
          </View>
          <View style={styles.textContainer}>
            <Text style={styles.name}>{this.props.chat.name}</Text>
            <View style={styles.circleAndDateContainer}>
              <Text style={styles.recentCircleName}>
                {this.props.chat.recentCircle}
              </Text>
              <Text style={styles.date}>{date.format('HH:MM A')}</Text>
            </View>
          </View>
        </TouchableOpacity>
        {this.renderCircleWindow()}
      </View>
    );
  }
}

export default ButtonView;
