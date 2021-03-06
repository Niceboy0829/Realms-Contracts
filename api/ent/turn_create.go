// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dopedao/RYO/api/ent/schema"
	"github.com/dopedao/RYO/api/ent/turn"
)

// TurnCreate is the builder for creating a Turn entity.
type TurnCreate struct {
	config
	mutation *TurnMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (tc *TurnCreate) SetUserID(s string) *TurnCreate {
	tc.mutation.SetUserID(s)
	return tc
}

// SetLocationID sets the "location_id" field.
func (tc *TurnCreate) SetLocationID(s string) *TurnCreate {
	tc.mutation.SetLocationID(s)
	return tc
}

// SetItemID sets the "item_id" field.
func (tc *TurnCreate) SetItemID(s string) *TurnCreate {
	tc.mutation.SetItemID(s)
	return tc
}

// SetBuyOrSell sets the "buy_or_sell" field.
func (tc *TurnCreate) SetBuyOrSell(b bool) *TurnCreate {
	tc.mutation.SetBuyOrSell(b)
	return tc
}

// SetAmountToGive sets the "amount_to_give" field.
func (tc *TurnCreate) SetAmountToGive(si schema.BigInt) *TurnCreate {
	tc.mutation.SetAmountToGive(si)
	return tc
}

// SetUserCombatStats sets the "user_combat_stats" field.
func (tc *TurnCreate) SetUserCombatStats(i []int) *TurnCreate {
	tc.mutation.SetUserCombatStats(i)
	return tc
}

// SetDrugLordCombatStats sets the "drug_lord_combat_stats" field.
func (tc *TurnCreate) SetDrugLordCombatStats(i []int) *TurnCreate {
	tc.mutation.SetDrugLordCombatStats(i)
	return tc
}

// SetTradeOccurs sets the "trade_occurs" field.
func (tc *TurnCreate) SetTradeOccurs(b bool) *TurnCreate {
	tc.mutation.SetTradeOccurs(b)
	return tc
}

// SetUserPreTradeItem sets the "user_pre_trade_item" field.
func (tc *TurnCreate) SetUserPreTradeItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPreTradeItem(si)
	return tc
}

// SetUserPostTradePreEventItem sets the "user_post_trade_pre_event_item" field.
func (tc *TurnCreate) SetUserPostTradePreEventItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPostTradePreEventItem(si)
	return tc
}

// SetUserPostTradePostEventItem sets the "user_post_trade_post_event_item" field.
func (tc *TurnCreate) SetUserPostTradePostEventItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPostTradePostEventItem(si)
	return tc
}

// SetUserPreTradeMoney sets the "user_pre_trade_money" field.
func (tc *TurnCreate) SetUserPreTradeMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPreTradeMoney(si)
	return tc
}

// SetUserPostTradePreEventMoney sets the "user_post_trade_pre_event_money" field.
func (tc *TurnCreate) SetUserPostTradePreEventMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPostTradePreEventMoney(si)
	return tc
}

// SetUserPostTradePostEventMoney sets the "user_post_trade_post_event_money" field.
func (tc *TurnCreate) SetUserPostTradePostEventMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetUserPostTradePostEventMoney(si)
	return tc
}

// SetMarketPreTradeItem sets the "market_pre_trade_item" field.
func (tc *TurnCreate) SetMarketPreTradeItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPreTradeItem(si)
	return tc
}

// SetMarketPostTradePreEventItem sets the "market_post_trade_pre_event_item" field.
func (tc *TurnCreate) SetMarketPostTradePreEventItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPostTradePreEventItem(si)
	return tc
}

// SetMarketPostTradePostEventItem sets the "market_post_trade_post_event_item" field.
func (tc *TurnCreate) SetMarketPostTradePostEventItem(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPostTradePostEventItem(si)
	return tc
}

// SetMarketPreTradeMoney sets the "market_pre_tradeMoney" field.
func (tc *TurnCreate) SetMarketPreTradeMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPreTradeMoney(si)
	return tc
}

// SetMarketPostTradePreEventMoney sets the "market_post_trade_pre_eventMoney" field.
func (tc *TurnCreate) SetMarketPostTradePreEventMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPostTradePreEventMoney(si)
	return tc
}

// SetMarketPostTradePostEventMoney sets the "market_post_trade_post_eventMoney" field.
func (tc *TurnCreate) SetMarketPostTradePostEventMoney(si schema.BigInt) *TurnCreate {
	tc.mutation.SetMarketPostTradePostEventMoney(si)
	return tc
}

// SetDealerDash sets the "dealer_dash" field.
func (tc *TurnCreate) SetDealerDash(b bool) *TurnCreate {
	tc.mutation.SetDealerDash(b)
	return tc
}

// SetWrangleDashedDealer sets the "wrangle_dashed_dealer" field.
func (tc *TurnCreate) SetWrangleDashedDealer(b bool) *TurnCreate {
	tc.mutation.SetWrangleDashedDealer(b)
	return tc
}

// SetMugging sets the "mugging" field.
func (tc *TurnCreate) SetMugging(b bool) *TurnCreate {
	tc.mutation.SetMugging(b)
	return tc
}

// SetRunFromMugging sets the "run_from_mugging" field.
func (tc *TurnCreate) SetRunFromMugging(b bool) *TurnCreate {
	tc.mutation.SetRunFromMugging(b)
	return tc
}

// SetGangWar sets the "gang_war" field.
func (tc *TurnCreate) SetGangWar(b bool) *TurnCreate {
	tc.mutation.SetGangWar(b)
	return tc
}

// SetDefendGangWar sets the "defend_gang_war" field.
func (tc *TurnCreate) SetDefendGangWar(b bool) *TurnCreate {
	tc.mutation.SetDefendGangWar(b)
	return tc
}

// SetCopRaid sets the "cop_raid" field.
func (tc *TurnCreate) SetCopRaid(b bool) *TurnCreate {
	tc.mutation.SetCopRaid(b)
	return tc
}

// SetBribeCops sets the "bribe_cops" field.
func (tc *TurnCreate) SetBribeCops(b bool) *TurnCreate {
	tc.mutation.SetBribeCops(b)
	return tc
}

// SetFindItem sets the "find_item" field.
func (tc *TurnCreate) SetFindItem(b bool) *TurnCreate {
	tc.mutation.SetFindItem(b)
	return tc
}

// SetLocalShipment sets the "local_shipment" field.
func (tc *TurnCreate) SetLocalShipment(b bool) *TurnCreate {
	tc.mutation.SetLocalShipment(b)
	return tc
}

// SetWarehouseSeizure sets the "warehouse_seizure" field.
func (tc *TurnCreate) SetWarehouseSeizure(b bool) *TurnCreate {
	tc.mutation.SetWarehouseSeizure(b)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TurnCreate) SetCreatedAt(t time.Time) *TurnCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TurnCreate) SetNillableCreatedAt(t *time.Time) *TurnCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// Mutation returns the TurnMutation object of the builder.
func (tc *TurnCreate) Mutation() *TurnMutation {
	return tc.mutation
}

// Save creates the Turn in the database.
func (tc *TurnCreate) Save(ctx context.Context) (*Turn, error) {
	var (
		err  error
		node *Turn
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TurnMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TurnCreate) SaveX(ctx context.Context) *Turn {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TurnCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TurnCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TurnCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := turn.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TurnCreate) check() error {
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "user_id"`)}
	}
	if v, ok := tc.mutation.UserID(); ok {
		if err := turn.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "user_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.LocationID(); !ok {
		return &ValidationError{Name: "location_id", err: errors.New(`ent: missing required field "location_id"`)}
	}
	if v, ok := tc.mutation.LocationID(); ok {
		if err := turn.LocationIDValidator(v); err != nil {
			return &ValidationError{Name: "location_id", err: fmt.Errorf(`ent: validator failed for field "location_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ItemID(); !ok {
		return &ValidationError{Name: "item_id", err: errors.New(`ent: missing required field "item_id"`)}
	}
	if v, ok := tc.mutation.ItemID(); ok {
		if err := turn.ItemIDValidator(v); err != nil {
			return &ValidationError{Name: "item_id", err: fmt.Errorf(`ent: validator failed for field "item_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.BuyOrSell(); !ok {
		return &ValidationError{Name: "buy_or_sell", err: errors.New(`ent: missing required field "buy_or_sell"`)}
	}
	if _, ok := tc.mutation.AmountToGive(); !ok {
		return &ValidationError{Name: "amount_to_give", err: errors.New(`ent: missing required field "amount_to_give"`)}
	}
	if _, ok := tc.mutation.UserCombatStats(); !ok {
		return &ValidationError{Name: "user_combat_stats", err: errors.New(`ent: missing required field "user_combat_stats"`)}
	}
	if _, ok := tc.mutation.DrugLordCombatStats(); !ok {
		return &ValidationError{Name: "drug_lord_combat_stats", err: errors.New(`ent: missing required field "drug_lord_combat_stats"`)}
	}
	if _, ok := tc.mutation.TradeOccurs(); !ok {
		return &ValidationError{Name: "trade_occurs", err: errors.New(`ent: missing required field "trade_occurs"`)}
	}
	if _, ok := tc.mutation.UserPreTradeItem(); !ok {
		return &ValidationError{Name: "user_pre_trade_item", err: errors.New(`ent: missing required field "user_pre_trade_item"`)}
	}
	if _, ok := tc.mutation.UserPostTradePreEventItem(); !ok {
		return &ValidationError{Name: "user_post_trade_pre_event_item", err: errors.New(`ent: missing required field "user_post_trade_pre_event_item"`)}
	}
	if _, ok := tc.mutation.UserPostTradePostEventItem(); !ok {
		return &ValidationError{Name: "user_post_trade_post_event_item", err: errors.New(`ent: missing required field "user_post_trade_post_event_item"`)}
	}
	if _, ok := tc.mutation.UserPreTradeMoney(); !ok {
		return &ValidationError{Name: "user_pre_trade_money", err: errors.New(`ent: missing required field "user_pre_trade_money"`)}
	}
	if _, ok := tc.mutation.UserPostTradePreEventMoney(); !ok {
		return &ValidationError{Name: "user_post_trade_pre_event_money", err: errors.New(`ent: missing required field "user_post_trade_pre_event_money"`)}
	}
	if _, ok := tc.mutation.UserPostTradePostEventMoney(); !ok {
		return &ValidationError{Name: "user_post_trade_post_event_money", err: errors.New(`ent: missing required field "user_post_trade_post_event_money"`)}
	}
	if _, ok := tc.mutation.MarketPreTradeItem(); !ok {
		return &ValidationError{Name: "market_pre_trade_item", err: errors.New(`ent: missing required field "market_pre_trade_item"`)}
	}
	if _, ok := tc.mutation.MarketPostTradePreEventItem(); !ok {
		return &ValidationError{Name: "market_post_trade_pre_event_item", err: errors.New(`ent: missing required field "market_post_trade_pre_event_item"`)}
	}
	if _, ok := tc.mutation.MarketPostTradePostEventItem(); !ok {
		return &ValidationError{Name: "market_post_trade_post_event_item", err: errors.New(`ent: missing required field "market_post_trade_post_event_item"`)}
	}
	if _, ok := tc.mutation.MarketPreTradeMoney(); !ok {
		return &ValidationError{Name: "market_pre_tradeMoney", err: errors.New(`ent: missing required field "market_pre_tradeMoney"`)}
	}
	if _, ok := tc.mutation.MarketPostTradePreEventMoney(); !ok {
		return &ValidationError{Name: "market_post_trade_pre_eventMoney", err: errors.New(`ent: missing required field "market_post_trade_pre_eventMoney"`)}
	}
	if _, ok := tc.mutation.MarketPostTradePostEventMoney(); !ok {
		return &ValidationError{Name: "market_post_trade_post_eventMoney", err: errors.New(`ent: missing required field "market_post_trade_post_eventMoney"`)}
	}
	if _, ok := tc.mutation.DealerDash(); !ok {
		return &ValidationError{Name: "dealer_dash", err: errors.New(`ent: missing required field "dealer_dash"`)}
	}
	if _, ok := tc.mutation.WrangleDashedDealer(); !ok {
		return &ValidationError{Name: "wrangle_dashed_dealer", err: errors.New(`ent: missing required field "wrangle_dashed_dealer"`)}
	}
	if _, ok := tc.mutation.Mugging(); !ok {
		return &ValidationError{Name: "mugging", err: errors.New(`ent: missing required field "mugging"`)}
	}
	if _, ok := tc.mutation.RunFromMugging(); !ok {
		return &ValidationError{Name: "run_from_mugging", err: errors.New(`ent: missing required field "run_from_mugging"`)}
	}
	if _, ok := tc.mutation.GangWar(); !ok {
		return &ValidationError{Name: "gang_war", err: errors.New(`ent: missing required field "gang_war"`)}
	}
	if _, ok := tc.mutation.DefendGangWar(); !ok {
		return &ValidationError{Name: "defend_gang_war", err: errors.New(`ent: missing required field "defend_gang_war"`)}
	}
	if _, ok := tc.mutation.CopRaid(); !ok {
		return &ValidationError{Name: "cop_raid", err: errors.New(`ent: missing required field "cop_raid"`)}
	}
	if _, ok := tc.mutation.BribeCops(); !ok {
		return &ValidationError{Name: "bribe_cops", err: errors.New(`ent: missing required field "bribe_cops"`)}
	}
	if _, ok := tc.mutation.FindItem(); !ok {
		return &ValidationError{Name: "find_item", err: errors.New(`ent: missing required field "find_item"`)}
	}
	if _, ok := tc.mutation.LocalShipment(); !ok {
		return &ValidationError{Name: "local_shipment", err: errors.New(`ent: missing required field "local_shipment"`)}
	}
	if _, ok := tc.mutation.WarehouseSeizure(); !ok {
		return &ValidationError{Name: "warehouse_seizure", err: errors.New(`ent: missing required field "warehouse_seizure"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "created_at"`)}
	}
	return nil
}

func (tc *TurnCreate) sqlSave(ctx context.Context) (*Turn, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TurnCreate) createSpec() (*Turn, *sqlgraph.CreateSpec) {
	var (
		_node = &Turn{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: turn.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: turn.FieldID,
			},
		}
	)
	if value, ok := tc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: turn.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := tc.mutation.LocationID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: turn.FieldLocationID,
		})
		_node.LocationID = value
	}
	if value, ok := tc.mutation.ItemID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: turn.FieldItemID,
		})
		_node.ItemID = value
	}
	if value, ok := tc.mutation.BuyOrSell(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldBuyOrSell,
		})
		_node.BuyOrSell = value
	}
	if value, ok := tc.mutation.AmountToGive(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldAmountToGive,
		})
		_node.AmountToGive = value
	}
	if value, ok := tc.mutation.UserCombatStats(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: turn.FieldUserCombatStats,
		})
		_node.UserCombatStats = value
	}
	if value, ok := tc.mutation.DrugLordCombatStats(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: turn.FieldDrugLordCombatStats,
		})
		_node.DrugLordCombatStats = value
	}
	if value, ok := tc.mutation.TradeOccurs(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldTradeOccurs,
		})
		_node.TradeOccurs = value
	}
	if value, ok := tc.mutation.UserPreTradeItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPreTradeItem,
		})
		_node.UserPreTradeItem = value
	}
	if value, ok := tc.mutation.UserPostTradePreEventItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPostTradePreEventItem,
		})
		_node.UserPostTradePreEventItem = value
	}
	if value, ok := tc.mutation.UserPostTradePostEventItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPostTradePostEventItem,
		})
		_node.UserPostTradePostEventItem = value
	}
	if value, ok := tc.mutation.UserPreTradeMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPreTradeMoney,
		})
		_node.UserPreTradeMoney = value
	}
	if value, ok := tc.mutation.UserPostTradePreEventMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPostTradePreEventMoney,
		})
		_node.UserPostTradePreEventMoney = value
	}
	if value, ok := tc.mutation.UserPostTradePostEventMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldUserPostTradePostEventMoney,
		})
		_node.UserPostTradePostEventMoney = value
	}
	if value, ok := tc.mutation.MarketPreTradeItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPreTradeItem,
		})
		_node.MarketPreTradeItem = value
	}
	if value, ok := tc.mutation.MarketPostTradePreEventItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPostTradePreEventItem,
		})
		_node.MarketPostTradePreEventItem = value
	}
	if value, ok := tc.mutation.MarketPostTradePostEventItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPostTradePostEventItem,
		})
		_node.MarketPostTradePostEventItem = value
	}
	if value, ok := tc.mutation.MarketPreTradeMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPreTradeMoney,
		})
		_node.MarketPreTradeMoney = value
	}
	if value, ok := tc.mutation.MarketPostTradePreEventMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPostTradePreEventMoney,
		})
		_node.MarketPostTradePreEventMoney = value
	}
	if value, ok := tc.mutation.MarketPostTradePostEventMoney(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: turn.FieldMarketPostTradePostEventMoney,
		})
		_node.MarketPostTradePostEventMoney = value
	}
	if value, ok := tc.mutation.DealerDash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldDealerDash,
		})
		_node.DealerDash = value
	}
	if value, ok := tc.mutation.WrangleDashedDealer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldWrangleDashedDealer,
		})
		_node.WrangleDashedDealer = value
	}
	if value, ok := tc.mutation.Mugging(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldMugging,
		})
		_node.Mugging = value
	}
	if value, ok := tc.mutation.RunFromMugging(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldRunFromMugging,
		})
		_node.RunFromMugging = value
	}
	if value, ok := tc.mutation.GangWar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldGangWar,
		})
		_node.GangWar = value
	}
	if value, ok := tc.mutation.DefendGangWar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldDefendGangWar,
		})
		_node.DefendGangWar = value
	}
	if value, ok := tc.mutation.CopRaid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldCopRaid,
		})
		_node.CopRaid = value
	}
	if value, ok := tc.mutation.BribeCops(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldBribeCops,
		})
		_node.BribeCops = value
	}
	if value, ok := tc.mutation.FindItem(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldFindItem,
		})
		_node.FindItem = value
	}
	if value, ok := tc.mutation.LocalShipment(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldLocalShipment,
		})
		_node.LocalShipment = value
	}
	if value, ok := tc.mutation.WarehouseSeizure(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: turn.FieldWarehouseSeizure,
		})
		_node.WarehouseSeizure = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: turn.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	return _node, _spec
}

// TurnCreateBulk is the builder for creating many Turn entities in bulk.
type TurnCreateBulk struct {
	config
	builders []*TurnCreate
}

// Save creates the Turn entities in the database.
func (tcb *TurnCreateBulk) Save(ctx context.Context) ([]*Turn, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Turn, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TurnMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TurnCreateBulk) SaveX(ctx context.Context) []*Turn {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TurnCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TurnCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
