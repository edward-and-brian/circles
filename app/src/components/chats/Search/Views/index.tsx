import React, { PureComponent } from 'react';
import { View, TouchableOpacity, Image, TextInput } from 'react-native';
import { Images, Colors } from '../../../../themes/';
import styles from './styles';

export interface Props {
  message: string;
  onMessageChange(newMessage: string): void;
}

class ScrollListView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <View style={styles.searchContainer}>
          <TextInput
            style={styles.messageInput}
            value={this.props.message}
            onChangeText={this.props.onMessageChange}
            placeholder="Search"
            placeholderTextColor={Colors.darkGray}
          />
        </View>
        <TouchableOpacity style={styles.buttonContainer}>
          <Image
            style={styles.buttonImage}
            source={Images.createMessage}
            resizeMode="contain"
          />
        </TouchableOpacity>
      </View>
    );
  }
}

export default ScrollListView;
