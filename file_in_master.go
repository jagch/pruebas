package main

import "fmt"

func function() {
	$account->info .= sprintf(
                            '<a href="#edit-accounts-modal" data-toggle="modal" data-payment-method="%s" data-payment-method-type="cryptocurrencies" data-wallet="%s" data-crypto-currency="%s" data-user-account-id="%s" data-network="%s" class="btn u-btn-3d u-btn-bluegray g-mt-5 mr-2 btn-sm"><i class="hs-admin-pencil"></i> %s</a>',
                            $account->payment_method_id,
                            $account->data->wallet,
                            $account->data->cryptocurrency,
                            $account->id,
                            $network,
                            _i('Edit')
                        );
                        if (Gate::allows('access', Permissions::$disable_user_account)) {
                            $account->info .= sprintf(
                                '<button type="button" id="disable-account" data-route="%s" data-payment-method-account="%s" class="btn u-btn-3d u-btn-primary g-mt-5 btn-sm"><i class="hs-admin-trash"></i> %s</button>',
                                route('betpay.accounts.user.disable', [$account->id]),
                                $account->payment_method_id,
                                _i('Delete')
                            );
                        }
}
