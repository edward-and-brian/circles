import React, { PureComponent } from 'react';
import NavigationService from '../../navigation/NavigationService';
import ChatsScreenView from './Views';
class ChatsScreen extends PureComponent {
  onPressChat() {
    NavigationService.navigate('circle');
  }

  render() {
    return <ChatsScreenView onPressChat={this.onPressChat} />;
  }
}

export default ChatsScreen;
