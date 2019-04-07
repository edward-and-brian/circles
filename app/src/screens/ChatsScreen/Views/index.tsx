import React, { PureComponent } from 'react';
import { View } from 'react-native';
import ChatsNavbar from '../../../components/chats/Navbar';
import ChatScrollList from '../../../components/chats/ScrollList';
import Fixtures from '../../../utils/Fixtures';
import styles from './styles';

export interface Props {
  onPressChat(): void;
}

const chats = Fixtures.getChats();

class ChatsScreenView extends PureComponent<Props> {
  render() {
    return (
      <View style={styles.container}>
        <ChatsNavbar />
        <View style={styles.searchContainer} />
        <ChatScrollList chats={chats} onPressChat={this.props.onPressChat} />
      </View>
    );
  }
}

export default ChatsScreenView;
