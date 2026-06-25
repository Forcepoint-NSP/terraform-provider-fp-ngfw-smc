resource "smc_single_fw" "tf_single_fw" {
  comment = "Created by Terraform"
  log_server_ref = "http://mysmc:8082/7.4/elements/log_server/1441"
  name = "tf_single_fw_modem_interface"
  nodes {
    firewall_node {
      name = "myfwnode"
      nodeid = 1
    }
  }
  physical_interfaces {
    physical_interface {
      aggregate_mode = "none"
      cvi_mode = "none"
      dhcp_server_on_interface {
        default_lease_time = 7200
      }
      duplicate_address_detection = true
      include_prefix_info_option_flag = true
      interface_id = "0"
      interfaces {
        single_node_interface {
          address = "192.168.100.14"
          auth_request = false
          auth_request_source = false
          automatic_default_route = false
          backup_heartbeat = false
          backup_mgt = false
          domain_specific_dns_queries_source = false
          dynamic = false
          network_value = "192.168.100.0/24"
          nicid = "0"
          nodeid = 1
          outgoing = false
          pppoa = false
          pppoe = false
          primary_heartbeat = false
          primary_mgt = true
          relayed_by_dhcp = false
          reverse_connection = false
          vrrp = false
          vrrp_id = -1
          vrrp_priority = -1
        }
      }
      lldp_mode = "disabled"
      log_moderation {
        burst = 1000
        log_event = "antispoofing"
        rate = 100
      }
      log_moderation {
        burst = 20000
        log_event = "discard"
        rate = 5000
      }
      log_moderation {
        burst = 80000
        log_event = "allow"
        rate = 40000
      }
      managed_address_flag = false
      mtu = -1
      name = "Interface 0"
      other_configuration_flag = false
      override_engine_settings = false
      override_log_moderation_settings = false
      qos_limit = -1
      qos_mode = "no_qos"
      route_replies_back_mode = false
      router_advertisement = false
      set_autonomous_address_flag = true
      shared_interface = false
      syn_mode = "default"
      sync_parameter {
        full_sync_interval = 5000
        heartbeat_group_ip = "224.0.0.221"
        incr_sync_interval = 50
        statesync_group_ip = "224.0.0.222"
        sync_mode = "sync_all"
        sync_security = "sign"
      }
      virtual_engine_vlan_ok = false
    }
  }
  physical_interfaces {
    modem_interface {
      apn_interfaces {
        apn = "internet"
        dhcp_relay {
          enabled = false
          trusted_circuit = false
        }
        dhcp_server_on_interface {
          default_lease_time = 7200
        }
        duplicate_address_detection = true
        include_prefix_info_option_flag = false
        interface_id = "0.1"
        interfaces {
          single_node_interface {
            auth_request = false
            auth_request_source = false
            automatic_default_route = true
            backup_heartbeat = false
            backup_mgt = false
            domain_specific_dns_queries_source = false
            dynamic = true
            dynamic_index = 1
            nicid = "0.1"
            nodeid = 1
            outgoing = false
            pppoa = false
            pppoe = false
            primary_heartbeat = false
            primary_mgt = false
            relayed_by_dhcp = false
            reverse_connection = false
            vrrp = false
            vrrp_id = -1
            vrrp_priority = -1
          }
        }
        managed_address_flag = false
        mtu = -1
        name = "APN internet (0.1)"
        other_configuration_flag = false
        override_engine_settings = false
        override_log_moderation_settings = false
        qos_limit = -1
        qos_mode = "no_qos"
        router_advertisement = false
        set_autonomous_address_flag = false
        shared_interface = false
        syn_mode = "default"
      }
      dhcp_server_on_interface {
        default_lease_time = 7200
      }
      duplicate_address_detection = true
      include_prefix_info_option_flag = true
      interface_id = "0"
      log_moderation {
        burst = 1000
        log_event = "antispoofing"
        rate = 100
      }
      log_moderation {
        burst = 20000
        log_event = "discard"
        rate = 5000
      }
      log_moderation {
        burst = 80000
        log_event = "allow"
        rate = 40000
      }
      managed_address_flag = false
      modem_interface_type = "LTE"
      mtu = -1
      name = "Modem 0"
      other_configuration_flag = false
      override_engine_settings = false
      override_log_moderation_settings = false
      pin_code = "123456"
      qos_limit = -1
      qos_mode = "no_qos"
      router_advertisement = false
      set_autonomous_address_flag = true
      shared_interface = false
      syn_mode = "default"
    }
  }
}
