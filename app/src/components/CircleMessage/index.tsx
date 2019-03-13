import React, { PureComponent } from 'react';
import CircleMessageView from './Views';

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

class CircleMessage extends PureComponent<Props, State> {
  render() {
    return <CircleMessageView {...this.props} />;
  }
}

export default CircleMessage;
