import React from 'react';

export const toStringlinefeed = (msg: string): JSX.Element[] => {
  const texts = msg.split(/(\n)/).map((item, index) => {
    return (
      <React.Fragment key={index}>
        {item.match(/\n/) ? <br /> : item}
      </React.Fragment>
    );
  });
  return texts;
};
