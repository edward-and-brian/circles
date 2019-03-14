import React, { PureComponent } from 'react';
import { View, Text, ScrollView, TouchableOpacity } from 'react-native';
import styles from './styles';

export interface Props {
  onPressChat(): void;
}
class ChatsScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.navbarContainer} />
        <View style={styles.searchContainer} />
        <View style={styles.chatsContainer}>
          <ScrollView>
            {[1, 2, 3, 4, 5, 6, 7, 8, 9, 10].map(chat => (
              <TouchableOpacity
                style={styles.chatContainer}
                key={`chat${chat}`}
                onPress={this.props.onPressChat}
              >
                <Text style={styles.chatName}>{chat}</Text>
              </TouchableOpacity>
            ))}
            <View style={styles.spacer} />
          </ScrollView>
        </View>
      </View>
    );
  }
}

export default ChatsScreenView;
