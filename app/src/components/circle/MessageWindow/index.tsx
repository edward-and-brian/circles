import React, { PureComponent } from 'react';
import MessageWindowView from './Views';
import Fixtures from '../../../utils/Fixtures';

const conversation = Fixtures.getTestMessages();

export interface Props {}
interface State {}

class MessageWindow extends PureComponent<Props, State> {
  render() {
    return <MessageWindowView conversation={conversation} />;
  }
}

export default MessageWindow;
