package client

import (
	"errors"

	"github.com/spectrocloud/hapi/models"
)

// remove models after refactoring to hapi models.
type VirtualMachine struct {
	APIVersion string   `json:"apiVersion"`
	APIGroup   string   `json:"apiGroup"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
type Spec struct {
	Status       string       `json:"status"`
	SpecTemplate SpecTemplate `json:"template"`
}

type SpecTemplate struct {
	Domain   Domain    `json:"domain"`
	Networks []Network `json:"networks"`
	Volumes  []Volume  `json:"volumes"`
}

type Domain struct {
	CPU       CPU       `json:"cpu"`
	Devices   Devices   `json:"devices"`
	Machine   Machine   `json:"machine"`
	Resources Resources `json:"resources"`
}

type CPU struct {
	Cores int `json:"cores"`
}

type Devices struct {
	Disks      []Disk      `json:"disks"`
	Interfaces []Interface `json:"interfaces"`
}

type Disk struct {
	Name     string   `json:"name"`
	DiskType DiskType `json:"disk"`
}

type DiskType struct {
	Bus string `json:"bus"`
}

type Interface struct {
	Masquerade Masquerade `json:"masquerade"`
	Name       string     `json:"name"`
}

type Masquerade struct{}

type Machine struct {
	Type string `json:"type"`
}

type Resources struct {
	Requests Requests `json:"requests"`
}

type Requests struct {
	Memory string `json:"memory"`
}

type Network struct {
	Name string `json:"name"`
	Pod  Pod    `json:"pod"`
}

type Pod struct{}

type Volume struct {
	Name             string           `json:"name"`
	ContainerDisk    ContainerDisk    `json:"containerDisk"`
	CloudInitNoCloud CloudInitNoCloud `json:"cloudInitNoCloud"`
}

type ContainerDisk struct {
	Image string `json:"image"`
}

type CloudInitNoCloud struct {
	UserData string `json:"userData"`
}

func (h *V1Client) GetVirtualMachines(cluster *models.V1SpectroCluster) error {
	return errors.New("not implemented")
}

func (h *V1Client) GetVirtualMachine(uid string) (VirtualMachine, error) {
	return VirtualMachine{}, errors.New("not implemented")
}

func (h *V1Client) UpdateVirtualMachine(cluster *models.V1SpectroCluster, body *VirtualMachine) error {
	return errors.New("not implemented")
}

func (h *V1Client) CreateVirtualMachine(uid string, body *VirtualMachine) error {
	return errors.New("not implemented")
}

func (h *V1Client) DeleteVirtualMachine(uid string, body *VirtualMachine) error {
	return errors.New("not implemented")
}
