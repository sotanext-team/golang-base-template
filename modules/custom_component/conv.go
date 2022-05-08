package custom_component

import (
    "app-api/ent"
    pb "github.com/es-hs/erpc/component"
)

func ToGRPCObject(component *ent.CustomComponent) (*pb.CustomComponent, error) {
    // Props must be preloaded
    props, err := component.Edges.PropsOrErr()
    if err != nil {
        return nil, err
    }
    var p []*pb.ComponentProperty
    for _, prop := range props {
        p = append(p, &pb.ComponentProperty{
            Id:    prop.ID,
            Name:  prop.Name,
            Value: prop.Value,
        })
    }
    return &pb.CustomComponent{
        Id:            component.ID,
        Name:          component.Name,
        Handle:        component.Handle,
        Content:       component.Content,
        EntryFileName: component.EntryFileName,
        Properties:    p,
    }, nil
}
