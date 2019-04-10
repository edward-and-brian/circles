import React, { PureComponent } from 'react';
import ScrollListView from './Views';
import moment from 'moment';

export interface Circle {
  name: string;
  avatar: string;
  lastMessage: {
    content: string;
    createdAt: moment.Moment;
  };
}

export interface Chat {
  name: string;
  avatar: string;
  circles: Circle[];
}

export interface Props {
  chats: Chat[];
  onPressChat(): void;
}

class ScrollList extends PureComponent<Props> {
  render() {
    return <ScrollListView {...this.props} />;
  }
}

export default ScrollList;
