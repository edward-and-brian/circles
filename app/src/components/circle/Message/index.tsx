import React, { PureComponent } from 'react';
import MessageView from './Views';

export interface Props {
  messageGroup: {
    user: {
      name: string;
      avatar: string;
    };
    messages: string[];
    date: Date;
  };
}

interface State {}

class Message extends PureComponent<Props, State> {
  render() {
    return <MessageView {...this.props} />;
  }
}

export default Message;
