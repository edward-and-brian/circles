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
      <Animated.ScrollView
        style={[
          styles.circleContainer,
          { maxHeight: this.props.circleWindowHeight },
        ]}
      >
        {this.props.chat.circles.map(circle => (
          <TouchableOpacity
            onPress={this.props.onPressChat}
            key={`circle${circle}`}
            style={styles.circleButton}
          >
            <View style={styles.avatarContainer}>
              <AvatarButton
                color={Colors.yellow}
                diameter={Scaled.screen.height * 0.04}
              />
            </View>
            <View style={styles.circleTextContainer}>
              <Text style={styles.circleName}>{this.props.chat.name}</Text>
              <View style={styles.recentMessageAndDateContainer}>
                <Text style={styles.recentMessageText}>
                  {this.props.chat.recentCircle}
                </Text>
                <Text style={styles.recentMessageDate}>
                  {date.format('h:mm a')}
                </Text>
              </View>
            </View>
          </TouchableOpacity>
        ))}
      </Animated.ScrollView>
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
              <Text style={styles.date}>{date.format('h:mm a')}</Text>
            </View>
          </View>
        </TouchableOpacity>
        {this.renderCircleWindow()}
      </View>
    );
  }
}

export default ButtonView;
