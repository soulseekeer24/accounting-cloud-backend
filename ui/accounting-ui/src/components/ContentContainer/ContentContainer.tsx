import React from 'react';

const ContentContainer = (props: React.PropsWithChildren<{}>) => (
    <div className="container">{props.children}</div>
)

export default ContentContainer;

