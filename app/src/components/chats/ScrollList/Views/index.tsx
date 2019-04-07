import React, { PureComponent } from 'react';
import { View, ScrollView } from 'react-native';
import { Chat } from '../';
import Button from '../../Button';
import styles from './styles';

export interface Props {
  chats: Chat[];
  onPressChat(): void;
}

class ScrollListView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <ScrollView>
          {this.props.chats.map(chat => (
            <Button
              chat={chat}
              onPressChat={this.props.onPressChat}
              key={`chat${chat.name}`}
            />
          ))}
        </ScrollView>
      </View>
    );
  }
}

export default ScrollListView;
