import React from 'react';

import { useUnraidRoute } from '~/state/unraid';
import { Description as SelectDescription } from './select/description';
import { Description as PlanDescription } from './plan/description';
import { Description as TransferDescription } from './transfer/description';
import { Line } from '~/shared/transfer/line';

export const Ticker: React.FunctionComponent = () => {
  const route = useUnraidRoute();

  return (
    <div>
      {(route === '/gather/select' || route === '/gather') && (
        <SelectDescription />
      )}
      {route === '/gather/plan' && <PlanDescription />}
      {route === '/gather/transfer/targets' && <TransferDescription />}
      {route === '/gather/transfer/operation' && <Line />}
    </div>
  );
};
