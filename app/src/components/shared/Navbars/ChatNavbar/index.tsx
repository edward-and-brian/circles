import React, { PureComponent } from 'react';
import ChatsNavbarView from './Views';

export interface Props {}

class ChatsNavbar extends PureComponent<Props> {
  render() {
    return <ChatsNavbarView {...this.props} />;
  }
}

export default ChatsNavbar;
